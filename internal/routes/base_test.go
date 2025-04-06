package routes_test

import (
	testhelpers "sandbox/internal/test_helpers"
	"testing"
)

func TestCheckPing(t *testing.T) {
	// given
	container := testhelpers.GetClean(t)
	srv := testhelpers.NewTestServer(t, container)

	t.Run("200 on /", func(t *testing.T) {
		// when, then
		srv.Get(t, "/").RequireOk(t)
	})
	t.Run("404 on unknown path", func(t *testing.T) {
		// when, then
		srv.Get(t, "/unknown").RequireStatus(t, 404)
	})

	t.Run("check secret", func(t *testing.T) {
		var res struct {
			Token string `json:"token"`
		}
		srv.Get(t, "/api/get_token").RequireOk(t).RequireUnmarshal(t, &res)
		t.Log(res.Token)

		srv.GetWithHeaders(t, "/api/secret/all_messages", map[string]string{"Authorization": res.Token}).RequireOk(t)
	})
}

package entities

type Pagination struct {
	PerPage int `query:"per_page"`
	Page    int `query:"page"`
}

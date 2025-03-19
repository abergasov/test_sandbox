package entities

type Pagination struct {
	PerPage int `json:"per_page"`
	Page    int `json:"page"`
}

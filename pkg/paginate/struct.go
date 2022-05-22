package paginate

type Paginate struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`

	BaseURL string      `json:"base_url"`
	Items   interface{} `json:"items"`

	OnEachSide int     `json:"on_each_side"`
	Fragment   *string `json:"fragment"`
}

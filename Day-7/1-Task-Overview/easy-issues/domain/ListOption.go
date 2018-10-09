package domain

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	Limit  int
	Offset int
}

// ListResponse
type ListResponse struct {
	data []interface{} `json:"data"`

	Page    int      `json:"page"`
	PerPage int      `json:"per_page"`
	Pages   int      `json:"pages"`
	Total   int      `json:"total"`
	Links   map[string]string `json:"links"`
}

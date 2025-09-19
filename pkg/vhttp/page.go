package vhttp

import "math"


const (
	defaultLimit       = 10
	defaultCurrentPage = 1
)
type Page struct {
	Total       int `json:"total"`
	Count       int `json:"count"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

func NewPaginator(totalCount, currentPage, limit int) *Page {
	pagination := Page{
		Total:       totalCount,
		CurrentPage: currentPage,
		PerPage:     limit,
	}
	if pagination.PerPage <= 0 {
		pagination.PerPage = defaultLimit
	}
	if pagination.CurrentPage <= 0 {
		pagination.CurrentPage = defaultCurrentPage
	}
	pagination.TotalPages = div(totalCount, limit)
	if pagination.CurrentPage < pagination.TotalPages {
		pagination.Count = pagination.PerPage
	} else if pagination.CurrentPage > pagination.TotalPages {
		pagination.Count = 0
	} else {
		pagination.Count = pagination.Total - ((pagination.CurrentPage - 1) * pagination.PerPage)
	}
	return &pagination
}

func div(a, b int) int {
	if b == 0 || a == 0 {
		return 0
	}
	return int(math.Ceil(float64(a) / float64(b)))
}

package common

// common.response
// Standard untuk menetapkan struktur response JSON
// yang dikirim ke client

// TO-DO : Kurang lebih sama kayak ErrorResponse
// cuman, ini tinggal implementasi bagian statusnya saja
type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

type SuccessResponse struct {
	// Status  int    `json:"status"`
	Message string          `json:"message"`
	Data    any             `json:"data,omitempty"`
	Meta    *PaginationMeta `json:"meta,omitempty"`
}

// TO-DO : perjelas response Error nanti
// perjelas kenapa error bisa terjadi, status code,
// dan bagian / field mana yang terjadi error
type ErrorResponse struct {
	// Status int `json:"status"`
	Message string `json:"message"`
	// Errors any `json:"errors,omitempty"`
}

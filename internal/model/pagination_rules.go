package model

// ---- aturan buat membatas pengiriman data dari database (pagination)

type PaginationParams struct {
	Page  int
	Limit int
}

func (p PaginationParams) Offset() int { return (p.Page - 1) * p.Limit }

func (p *PaginationParams) Normalize() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = 10
	}
	if p.Limit > 30 {
		p.Limit = 30 // batas agar tidak membebani database
	}
}

type PaginatedResultData[T any] struct {
	Data       []T
	Page       int
	Limit      int
	TotalItems int
	TotalPages int
}

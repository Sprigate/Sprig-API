package expenses

import (
	"sprig/internal/model"
	"time"
)

type CreateExpenseRequest struct {
	UserID uint64 `json:"user_id"`
	// tadinya mau ada field reason untuk memperjelas pengeluaran
	// tapi kayaknya di tulis di deskripsi juga cukup
	CategoryID  uint64            `json:"category_id"`
	Amount      uint64            `json:"amount"`
	Description *string           `json:"description"` // pertimbangkan jadi field wajib, jadi gaada field reason
	ExpenseDate string            `json:"expense_date"`
	Type        model.ExpenseType `json:"type"`
	Wallet      model.Wallet      `json:"wallet"`
}

type CreateUpdateExpenseRequest struct {
	UserID      uint64             `json:"user_id"`
	CategoryID  *uint64            `json:"category_id"`
	Amount      *uint64            `json:"amount"`
	Description *string            `json:"description"`
	ExpenseDate *string            `json:"expense_date"`
	Type        *model.ExpenseType `json:"type"`
	Group       *model.Wallet      `json:"wallet"`
}

// buat dynamic filter di expense
type ExpenseFilterRequest struct {
	StartDate  string `form:"start_date" json:"start_date"`
	EndDate    string `form:"end_date" json:"end_date"`
	CategoryID string `form:"category_id" json:"category_id"`
	Type       string `form:"type" json:"type"`
	// pertimbangan buat nanti: pake wallet gak buat filtering?
}

type ExpenseResponse struct {
	ID          uint64            `json:"id"`
	UserID      uint64            `json:"user_id"`
	CategoryID  uint64            `json:"category_id"`
	Amount      uint64            `json:"amount"`
	Description string            `json:"description"`
	ExpenseDate time.Time         `json:"expense_date"`
	Type        model.ExpenseType `json:"type"`
	Wallet      model.Wallet      `json:"wallet"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

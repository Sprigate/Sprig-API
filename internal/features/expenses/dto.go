package expenses

import (
	"sprig/internal/model"
	"time"
)

type CreateExpenseRequest struct {
	UserID       uint64            `json:"user_id"`
	CategoryID   uint64            `json:"category_id"`
	Amount       uint64            `json:"amount"`
	Descripttion *string           `json:"description"`
	ExpenseDate  time.Time         `json:"expense_date"`
	Type         model.ExpenseType `json:"type"`
}

type ExpenseResponse struct {
	ID          uint64            `json:"id"`
	UserID      uint64            `json:"user_id"`
	CategoryID  uint64            `json:"category_id"`
	Amount      uint64            `json:"amount"`
	Description *string           `json:"description"`
	ExpenseDate time.Time         `json:"expense_date"`
	Type        model.ExpenseType `json:"type"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

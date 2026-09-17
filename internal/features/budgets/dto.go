package budgets

import "time"

type CreateBudgetRequest struct {
	UserID            uint64 `json:"user_id"`
	TotalIncome       uint64 `json:"total_income"`
	NeedsPercentage   uint8  `json:"needs_percentage"`
	WantsPercentage   uint8  `json:"wants_percentage"`
	SavingsPercentage uint8  `json:"savings_percentage"`
	Year              uint16 `json:"year"`
	Month             uint8  `json:"month"`
}

type CreateBudgetUpdateRequest struct {
	UserID            uint64  `json:"user_id"`
	TotalIncome       *uint64 `json:"total_income"`
	NeedsPercentage   *uint8  `json:"needs_percentage"`
	WantsPercentage   *uint8  `json:"wants_percentage"`
	SavingsPercentage *uint8  `json:"savings_percentage"`
	Year              *uint16 `json:"year"`
	Month             *uint8  `json:"month"`
}

type BudgetResponse struct {
	ID                uint64    `json:"id"`
	UserID            uint64    `json:"user_id"`
	TotalIncome       uint64    `json:"total_income"`
	NeedsPercentage   uint8     `json:"needs_percentage"`
	WantsPercentage   uint8     `json:"wants_percentage"`
	SavingsPercentage uint8     `json:"savings_percentage"`
	Year              uint16    `json:"year"`
	Month             uint8     `json:"month"`
	UpdatedAt         time.Time `json:"updated_at"`
}

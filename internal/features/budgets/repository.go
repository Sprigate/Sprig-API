package budgets

import "sprig/internal/model"

type BudgetsRepository interface {
	AddBudget(budget *model.Budget) error
	Save(budget *model.Budget) error
	DeleteByID(id uint64) error
	FindByID(id uint64) (*model.Budget, error)
	FindAll() ([]model.Budget, error)
	FilterByMonthAndYear(month uint8, year uint16) ([]model.Budget, error)
}

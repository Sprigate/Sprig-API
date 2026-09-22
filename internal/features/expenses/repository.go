package expenses

import "sprig/internal/model"

type ExpensesRepository interface {
	AddExpense(expenses *model.Expense) error
	Save(expense *model.Expense) error
	DeleteByID(id uint64) error
	FindAll() ([]model.Expense, error)
	FindByID(id uint64) (*model.Expense, error)
	FindByFilter(filter ExpenseFilterService) ([]model.Expense, error)
}

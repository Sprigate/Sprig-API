package bootstrap

import (
	"database/sql"
	"sprig/internal/features/budgets"
	"sprig/internal/features/categories"
	"sprig/internal/features/expenses"
	"sprig/internal/features/users"
)

// List layer yang mau diinject
type Container struct {
	UserHandler     *users.UserHandler
	CategoryHandler *categories.CategoryHandler
	BudgetHandler   *budgets.BudgetHandler
	ExpenseHandler  *expenses.ExpenseHandler
}

// Setiap ada DI, tambahnya disini aja
func NewContainer(db *sql.DB) *Container {
	usersRepository := users.NewUserRepository(db)
	usersService := users.NewUserService(usersRepository)
	usersHandler := users.NewUserHandler(usersService)

	categoriesRepository := categories.NewCategoryRepository(db)
	categoriesService := categories.NewCategoryService(categoriesRepository)
	categoriesHandler := categories.NewCategoryHandler(categoriesService)

	budgetsRepository := budgets.NewBudgetRepository(db)
	budgetsService := budgets.NewBudgetService(budgetsRepository)
	budgetsHandler := budgets.NewBudgetHandler(budgetsService)

	expenseRepository := expenses.NewExpenseRepository(db)
	expenseService := expenses.NewExpenseService(expenseRepository)
	expenseHandler := expenses.NewExpenseHandler(expenseService)

	return &Container{
		UserHandler:     usersHandler,
		CategoryHandler: categoriesHandler,
		BudgetHandler:   budgetsHandler,
		ExpenseHandler:  expenseHandler,
	}
}

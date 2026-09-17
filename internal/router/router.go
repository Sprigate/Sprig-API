package router

import (
	"net/http"
	"sprig/internal/bootstrap"
	"sprig/internal/features/budgets"
	"sprig/internal/features/categories"
	"sprig/internal/features/users"

	httpSwagger "github.com/swaggo/http-swagger"
)

func New(c *bootstrap.Container) *http.ServeMux {
	mux := http.NewServeMux()

	const apiV1 = "/api/v1"

	users.RegisterRoute(mux, apiV1, c.UserHandler)
	categories.RegisterRoute(mux, apiV1, c.CategoryHandler)
	budgets.RegisterRoute(mux, apiV1, c.BudgetHandler)

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	return mux
}

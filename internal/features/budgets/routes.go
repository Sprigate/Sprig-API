package budgets

import "net/http"

func RegisterRoute(mux *http.ServeMux, prefix string, h *BudgetHandler) {
	mux.HandleFunc("POST "+prefix+"/budgets", h.Create)
	mux.HandleFunc("PATCH "+prefix+"/budgets/{id}", h.Update)
	mux.HandleFunc("DELETE "+prefix+"/budgets/{id}", h.Delete)
	mux.HandleFunc("GET "+prefix+"/budgets", h.FindAll)
}

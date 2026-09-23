package expenses

import "net/http"

func RegisterRoute(mux *http.ServeMux, prefix string, h *ExpenseHandler) {
	mux.HandleFunc("POST "+prefix+"/expenses", h.Create)
	mux.HandleFunc("PATCH "+prefix+"/expenses/{id}", h.Update)
	mux.HandleFunc("DELETE "+prefix+"/expenses/{id}", h.Delete)
	mux.HandleFunc("GET "+prefix+"/expenses", h.FindAll)
}

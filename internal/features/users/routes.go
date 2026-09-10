package users

import "net/http"

// USERS ENDPOINT
func RegisterRoute(mux *http.ServeMux, prefix string, h *UserHandler) {
	mux.HandleFunc("POST "+prefix+"/users", h.Create)
	mux.HandleFunc("GET "+prefix+"/users", h.FindAll)
	mux.HandleFunc("PATCH "+prefix+"/users/{id}", h.Update)
	mux.HandleFunc("DELETE "+prefix+"/users/{id}", h.Delete)
}

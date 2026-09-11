package categories

import (
	"net/http"
)

func RegisterRoute(mux *http.ServeMux, prefix string, h *CategoryHandler) {
	mux.HandleFunc("POST "+prefix+"/categories", h.Create)
	mux.HandleFunc("GET "+prefix+"/categories", h.FindAll)
	mux.HandleFunc("GET "+prefix+"/categories/{name}", h.FindByName)
	mux.HandleFunc("PATCH "+prefix+"/categories/{id}", h.Update)
	mux.HandleFunc("DELETE "+prefix+"/categories/{id}", h.Delete)
}

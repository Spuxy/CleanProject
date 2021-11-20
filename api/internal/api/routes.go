package api

import (
	"github.com/gorilla/mux"
)

type Router struct {
	handler *Handler
}

func NewRouter(h *Handler) *Router {
	return &Router{h}
}

// TODO: Implementace middlewaru
func (r *Router) Routes() {
	mux := mux.NewRouter()
	mux.HandleFunc("/test", r.handler.Welcome)
}

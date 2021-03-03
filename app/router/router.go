package router

import (
	"github.com/go-chi/chi"
	"lcng/app/app"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.MethodFunc("GET", "/", app.HandleIndex)
	return r
}
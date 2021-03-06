package router

import (
	"github.com/go-chi/chi"
	"lcng/app/app"
	"lcng/app/requestlog"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))

	return r
}
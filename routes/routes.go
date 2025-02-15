package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/svuvi/goflashcards/assets"
	"github.com/svuvi/goflashcards/layouts"
)

type BaseHandler struct {
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (h *BaseHandler) NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", templ.Handler(layouts.Index()))
	mux.Handle("GET /static/", http.FileServer(http.FS(assets.Static)))

	mux.Handle("GET /feedback", templ.Handler(layouts.Feeback()))
	/* mux.Handle("GET /make")
	mux.Handle("GET /my")
	mux.Handle("GET /find") */

	return mux
}

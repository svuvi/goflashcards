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
	mux.Handle("/static/", http.FileServer(http.FS(assets.Static)))

	return mux
}

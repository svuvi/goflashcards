package routes

import (
	"net/http"

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

	mux.HandleFunc("GET /{$}", h.index)
	mux.Handle("/static/", http.FileServer(http.FS(assets.Static)))

	return mux
}

func (h *BaseHandler) index(w http.ResponseWriter, r *http.Request) {
	render(w, r, layouts.Index())
}

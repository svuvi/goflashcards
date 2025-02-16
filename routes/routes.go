package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"
	"github.com/svuvi/goflashcards/assets"
	"github.com/svuvi/goflashcards/layouts"
	"github.com/svuvi/goflashcards/models"
	"github.com/svuvi/goflashcards/repositories"
)

type BaseHandler struct {
	SetRepo  models.FlashcardSetRepository
	CardRepo models.CardRepository
}

func NewBaseHandler(db *sqlx.DB) *BaseHandler {
	return &BaseHandler{
		SetRepo:  repositories.NewFlashcardSetRepo(db),
		CardRepo: repositories.NewCardRepo(db),
	}
}

func (h *BaseHandler) NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", templ.Handler(layouts.Index()))
	mux.Handle("GET /static/", http.FileServer(http.FS(assets.Static)))

	mux.Handle("GET /feedback", templ.Handler(layouts.Feeback()))
	mux.Handle("GET /make", templ.Handler(layouts.Make()))
	mux.Handle("GET /my", templ.Handler(layouts.My()))
	mux.Handle("GET /find", templ.Handler(layouts.Find()))

	mux.HandleFunc("GET /set/{setID}/{slug}", h.setView)
	mux.HandleFunc("GET /set/{setID}/{slug}/{cardNumber}", h.setView)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		render(w, r, layouts.Err(http.StatusNotFound, http.StatusText(http.StatusNotFound)))
	})

	return mux
}

func (h *BaseHandler) setView(w http.ResponseWriter, r *http.Request) {
	setID, err := strconv.Atoi(r.PathValue("setID"))
	if err != nil {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	set, err := h.SetRepo.Get(setID)
	if err == sql.ErrNoRows || set.Slug != r.PathValue("slug") {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	cardNumber, err := strconv.Atoi(r.PathValue("cardNumber"))
	if err != nil {
		cardNumber = 1
	}

	cardsTotal, _ := h.CardRepo.CountCardsInSet(setID)
	card, _ := h.CardRepo.GetNthCard(setID, cardNumber)

	render(w, r, layouts.Set(layouts.SetProps{
		Set:            set,
		Card:           card,
		CardsTotal:     cardsTotal,
		ThisCardNumber: cardNumber,
	}))
}

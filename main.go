package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/svuvi/goflashcards/db"
	"github.com/svuvi/goflashcards/middleware"
	"github.com/svuvi/goflashcards/routes"
)

func main() {
	addr := flag.String("addr", ":7010", "Адрес сервера")

	db := db.ConnectDB("database.db")
	defer db.Close()

	h := routes.NewBaseHandler(db)
	router := middleware.NewLogger(h.NewRouter())

	server := http.Server{
		Addr:    *addr,
		Handler: router,
	}

	log.Printf("Сервер запущен на http://localhost%s", *addr)
	server.ListenAndServe()
}

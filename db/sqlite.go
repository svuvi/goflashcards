package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB(location string) *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", location)
	err := db.Ping()
	if err != nil {
		log.Fatal("Impossible to connect to the DB:\n", err)
	}

	return db
}

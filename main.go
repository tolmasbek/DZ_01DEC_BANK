package main

import (
	"bank-t/database"
	"bank-t/pkg/authorisation"
	"database/sql"
	"log"
	_ "mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "bank-t")
	if err != nil {
		log.Fatalf("Err is %e")
	}
	database.DbInit(db)

	Start(db)
}

func Start(db *sql.DB) {
	for {
		login, password, role := authorisation.Authorisation(db)
		authorisation.UserAdmin(db, login, password, role)
	}
}

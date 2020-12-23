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

func Start(datab *sql.DB) {
	for {
		login, password, role := authorisation.Authorisation(datab)
		authorisation.GetLoginPassUserOrAdmin(datab, login, password, role)
	}
}

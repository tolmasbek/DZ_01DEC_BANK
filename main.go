package main

import (
	"bank-t/database"
	"bank-t/pkg/authorisation"
	"bank-t/pkg/modules"
	"database/sql"
	"log"
	_"mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "bank-t")
	if err != nil {
		log.Fatalf("Err is %e")
	}
	database.DbInit(db)

	Start(db)
}

func Start(db *sql.DB){
	for{
		login, password := authorisation.Authorisation(db)
		modules.Login(db, login, password)
	}
}

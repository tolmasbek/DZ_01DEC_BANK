package database

import (
	"database/sql"
	"log"
)

func DbInit(db *sql.DB) {
	Tables := []string{CreateTableUsers, CreateTableAccounts, CreateTableCurrencies, CreateTableATMs}
	for _, ddl := range Tables {
		_, err := db.Exec(ddl)
		if err != nil {
			log.Fatalf("Can't init %s err is %e", ddl, err)
		}
	}
}

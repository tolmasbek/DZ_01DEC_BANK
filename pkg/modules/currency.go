package modules

import (
	"bank-t/database"
	"database/sql"
	"fmt"
)

type CurrencyB struct {
	Id   int64
	Name string
}

func AddCurr(dataBase *sql.DB, currency CurrencyB) (err error) {
	_, err = dataBase.Exec(database.AddNewCurrency, currency.Name)
	if err != nil {
		fmt.Println(err)
	}
	return
}

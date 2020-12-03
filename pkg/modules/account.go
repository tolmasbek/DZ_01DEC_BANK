package modules

import (
	"bank-t/database"
	"database/sql"
	"fmt"
)

type Account struct {
	Id        int64
	UserId    int64
	NumberAcc int64
	Amount    int64
	Currency  int64
	Remove    bool
}

func AddNewAccount(dataBase *sql.DB, accounts Account)(err error){
	_, err = dataBase.Exec(database.AddNewAcc, accounts.UserId, accounts.NumberAcc, accounts.Amount, accounts.Currency, accounts.Remove)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
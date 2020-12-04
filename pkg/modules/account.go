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
	Pin       int64
	Remove    bool
}

func AddNewAccount(dataBase *sql.DB, accounts Account) (err error) {
	_, err = dataBase.Exec(database.AddNewAcc, accounts.UserId, accounts.NumberAcc, accounts.Amount, accounts.Currency, accounts.Pin, accounts.Remove)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func ShowAcc(dataBase *sql.DB, pin int64) {
	var Acc Account
	_ = dataBase.QueryRow(`Select *From accounts Where pin=($1)`, pin).Scan(
		&Acc.Id,
		&Acc.UserId,
		&Acc.NumberAcc,
		&Acc.Amount,
		&Acc.Currency,
		&Acc.Pin,
		&Acc.Remove)
		if Acc.Pin == pin {
			fmt.Println(Acc)
		}
}

package modules

import (
	"bank-t/database"
	"database/sql"
	"fmt"
	"os"
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
	_, err = dataBase.Exec(database.AddNewAcc,
		accounts.UserId,
		accounts.NumberAcc,
		accounts.Amount,
		accounts.Currency,
		accounts.Pin,
		accounts.Remove)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func ShowAcc(dataBase *sql.DB, pin int64) {
	var Acc Account
	row := dataBase.QueryRow(database.ShowDataAcc, pin)
	err := row.Scan(
		&Acc.Id,
		&Acc.NumberAcc,
		&Acc.Amount,
		&Acc.Pin,
	)
	if err != nil {
		fmt.Println("Неверный PIN код: ", err)
	}
	if Acc.Pin == pin {
		fmt.Println(Acc)
	} else {
		fmt.Println("Введите заново PIN код")
	}
}

func ChekAccountNumber(dbase *sql.DB, num1 int64) {
	var numAcc Account
	err := dbase.QueryRow(database.ChekNumAcc, num1).Scan(&numAcc.NumberAcc)
	if numAcc.NumberAcc != num1 {
		fmt.Println("Неверный номер аккаунта ChekAccountNumber")
		os.Exit(0)
	}
	if err != nil {
		fmt.Println("Ошибка в AddTransaction", err)
	}
}

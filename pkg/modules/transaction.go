package modules

import (
	"bank-t/database"
	"database/sql"
	"fmt"
	"os"
)

type Transaction struct {
	Id               int64
	AccNumbSender    int64
	AccNumbAddressee int64
	TranslatedSum    int64
	Date             string
	Time             string
}

func AddNewTransaction(db *sql.DB, tran Transaction) (ok bool, err error) {
	_, err = db.Exec(database.AddNewTransaction,
		tran.AccNumbSender,
		tran.AccNumbAddressee,
		tran.TranslatedSum,
		tran.Date,
		tran.Time,
	)
	if err != nil {
		fmt.Println("Can't inserted", err)
		return false, err
	}
	return true, nil
}

func HistoryOfTransactions(db *sql.DB, numbAcc, pinCode int64) {
	var Acc Account
	row2 := db.QueryRow(database.ChekDataAcc, numbAcc, pinCode)
	err1 := row2.Scan(
		&Acc.NumberAcc,
		&Acc.Pin,
	)
	if (Acc.NumberAcc != numbAcc) && (Acc.Pin != pinCode) {
		fmt.Println("Неверный номер аккаунта или PIN - code", err1)
		os.Exit(0)
	}
	if err1 != nil {
		fmt.Println("Неверный PIN код: ", err1)
	}

	var tran Transaction
	row := db.QueryRow(database.ShowHistoryTran, numbAcc)
	err := row.Scan(
		&tran.Id,
		&tran.AccNumbSender,
		&tran.AccNumbAddressee,
		&tran.TranslatedSum,
		&tran.Date,
		&tran.Time,
	)
	if err != nil {
		fmt.Println("Ошибка в HistoryOfTransactions: ", err)
	}
	fmt.Println(tran)
}

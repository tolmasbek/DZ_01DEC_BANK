package authorisation

import (
	"bank-t/database"
	"bank-t/pkg/modules"
	"database/sql"
	"fmt"
	"os"
	"time"
)

func TransfersFromAccInAcc(dbase *sql.DB, sum, numAcc1, numAcc2 int64) {
	var trFromAccInAcc1 modules.Account
	err := dbase.QueryRow(database.TransferFromAcc, sum, numAcc1).Scan(
		&trFromAccInAcc1.NumberAcc,
		&trFromAccInAcc1.Amount,
	)
	if err != nil {
		fmt.Println("Ошибка в TransfersFromAcc", err)
	}

	err1 := dbase.QueryRow(database.TransferToAcc, sum, numAcc2).Scan(
		&trFromAccInAcc1.NumberAcc,
		&trFromAccInAcc1.Amount,
	)
	if err1 != nil {
		fmt.Println("Ошибка в TransfersInAcc", err1)
	}
}

func AddTransaction(db *sql.DB) {
	var accNumbSender, accNumbAddressee, translatedSum int64
	fmt.Println("Введите номер вашего аккаунта: ")
	fmt.Scan(&accNumbSender)

	var pinc int64
	fmt.Print("Введите PIN-код вашего счета: ")
	fmt.Scan(&pinc)
	ChekAccountNumber(db, accNumbSender, pinc)

	fmt.Println("Введите номер аккаунта получателя суммы: ")
	fmt.Scan(&accNumbAddressee)
	modules.ChekAccountNumber(db, accNumbAddressee)

	fmt.Println("Введите сумму: ")
	fmt.Scan(&translatedSum)
	ChekAccAmount(db, translatedSum)

	var now time.Time = time.Now()
	data := now.Format("02-Jan-2006")
	format := now.Format("15:04")

	newTran := modules.Transaction{
		Id:               0,
		AccNumbSender:    accNumbSender,
		AccNumbAddressee: accNumbAddressee,
		TranslatedSum:    translatedSum,
		Date:             data,
		Time:             format,
	}
	_, err := modules.AddNewTransaction(db, newTran)
	if err != nil {
		fmt.Println("Vse ploho")
		return
	}

	TransfersFromAccInAcc(db, translatedSum, accNumbSender, accNumbAddressee)

}

func ChekAccountNumber(dbase *sql.DB, numAc, Pin int64) {
	var numAcc modules.Account
	row2 := dbase.QueryRow(database.ChekDataAcc, numAc, Pin)
	err1 := row2.Scan(
		&numAcc.NumberAcc,
		&numAcc.Pin,
	)
	if (numAcc.NumberAcc != numAc) && (numAcc.Pin != Pin) {
		fmt.Println("Неверный номер аккаунта или PIN - code", err1)
		os.Exit(0)
	}
	if err1 != nil {
		fmt.Println("Неверный PIN код: ", err1)
	}
}

func ChekAccAmount(dbase *sql.DB, sum int64) {
	var CheckAmountAcc modules.Account
	err := dbase.QueryRow(database.ChekAmountAcc, sum).Scan(
		&CheckAmountAcc.Amount,
	)
	if err != nil {
		fmt.Println("Ошибка в AddTransaction", err)
	}
	if CheckAmountAcc.Amount <= sum {
		fmt.Println("У вас недостаточно средств для перевода !!!")
	}
}

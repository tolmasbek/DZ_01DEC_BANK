package authorisation

import (
	"bank-t/database"
	"bank-t/pkg/modules"
	"database/sql"
	"fmt"
	"time"
)

type Payment struct {
	Id         int64
	NameSrv    string
	AccNumb    int64
	SumRemoved int64
	SumFilled  int64
	Date       string
	TimePay       string
}

const ChoiceServices = `Наши услуги: 
	Пополнить счет вашего аккаунта
	Снятие со счета`

func SrvcWin(dab *sql.DB)  {
	fmt.Println(ChoiceServices)
	var nameSer, data, vremya string
	var svoySchet = "1. Пополнить свой счет"
	var snyatSchet = "2. Снятие со счета"
	var accountNum, snyatSum, popolnSum int64

	fmt.Println("Введите номер вашего аккаунта: ")
	fmt.Scan(&accountNum)

	var pinc int64
	fmt.Print("Введите PIN-код вашего счета: ")
	fmt.Scan(&pinc)
	ChekAccountNumber(dab, accountNum, pinc)
	var chServ int64
	fmt.Println("Выберите услугу: ")
	fmt.Println(svoySchet,"\n",snyatSchet)
	fmt.Scan(&chServ)
	if chServ == 1{
		nameSer = svoySchet
		fmt.Println("Введите сумму для пополнения:")
		fmt.Scan(&popolnSum)
		AddPayment(dab, popolnSum, accountNum)
	}
	if chServ == 2{
		nameSer = snyatSchet
		fmt.Println("Введите сумму для снятия:")
		fmt.Scan(&snyatSum)
		SubPayment(dab, snyatSum, accountNum)
	}

	var now time.Time = time.Now()
	data = now.Format("02-Jan-2006")
	vremya = now.Format("15:04")

	var srv Payment
	srv = Payment{
		Id:         0,
		NameSrv:    nameSer,
		AccNumb:    accountNum,
		SumRemoved: snyatSum,
		SumFilled:  popolnSum,
		Date:       data,
		TimePay:    vremya,
	}
	_, err := AddNewSrvc(dab, srv)
	if err != nil {
		fmt.Println("Vse ploho")
		return
	}
}

func AddNewSrvc(db *sql.DB, pay Payment) (ok bool, err error) {
	_, err = db.Exec(database.AddNewServices,
		pay.NameSrv,
		pay.AccNumb,
		pay.SumRemoved,
		pay.SumFilled,
		pay.Date,
		pay.TimePay,
	)
	if err != nil {
		fmt.Println("Can't inserted", err)
		return false, err
	}
	return true, nil
}

func AddPayment(db *sql.DB, summ, num int64)  {
	var addPay modules.Account
	err1 := db.QueryRow(database.TransferToAcc, summ, num).Scan(
		&addPay.NumberAcc,
		&addPay.Amount,
	)
	if err1 != nil {
		fmt.Println("Ошибка в AddPayment", err1)
	}
}

func SubPayment(db *sql.DB, summ1, num1 int64)  {
	var subPay modules.Account
	err := db.QueryRow(database.TransferFromAcc, summ1, num1).Scan(
		&subPay.NumberAcc,
		&subPay.Amount,
	)
	if err != nil {
		fmt.Println("Ошибка в SubPayment", err)
	}
}

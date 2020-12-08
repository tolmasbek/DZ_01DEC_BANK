package authorisation

import (
	"bank-t/pkg/modules"
	"database/sql"
	"fmt"
	"log"
	"os"
)

const FirstWindow = `"BANK-T"
 	1. Авторизация (Вход)
	0. Выход`

func Authorisation(db *sql.DB) (login, password, role string) {
	var firstChoice int64
	fmt.Println(FirstWindow)
	fmt.Scan(&firstChoice)
	switch firstChoice {
	case 1:
	m1:
		fmt.Println("Введите логин и пароль: ")
		fmt.Print("логин: ")
		fmt.Scan(&login)
		var pass string
		fmt.Print("пароль: ")
		fmt.Scan(&pass)
		fmt.Print("повторите пароль: ")
		fmt.Scan(&password)
		if pass == password {
			var role int64
			fmt.Print("Выберите ваш роль 1.admin | 2.user: ")
			fmt.Scan(&role)
			switch role {
			case 1:
				rAdmin := "admin"
				GetLoginPassUserOrAdmin(db, login, password, rAdmin)
			case 2:
				rUser := "user"
				GetLoginPassUserOrAdmin(db, login, password, rUser)
			default:
				fmt.Println("Неправильный ввод!!!")
			}
		} else {
			goto m1
		}
	case 0:
		fmt.Println("////////////////////////////////////")
		os.Exit(0)
	default:
		fmt.Println("Not correct data!!")
	}
	return
}

const AdminChoice = `Выберите дейсвтия:
	1. Добавить пользователя	2. Добавить валюту
	3. Добавить автомат		4. Открыть счет
	5. Посмотреть мои данные	6. Мои счета
	7. Переводы
	0. Выход`

func AdminWin(db *sql.DB, admin modules.User) {
	fmt.Println(AdminChoice)
	var adminCh int64
	fmt.Scan(&adminCh)
	switch adminCh {
	case 1:
		AddUser(db)
	case 2:
		AddCurrency(db)
	case 3:
		AddAtm(db)
	case 4:
		AddAccount(db)
	case 5:
		fmt.Println(admin)
	case 6:
		var pinAdm int64
		fmt.Print("Введите pin вашего счета: ")
		fmt.Scan(&pinAdm)
		modules.ShowAcc(db, pinAdm)
	case 7:
		TransferWin(db)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Не корректный ввод попробуйте еще раз")
	}
}

const UserChoice = `Выберите дейсвтия:
	1. Посмотреть мои данные	2. Открыть счет
	3. Мои счета			4. Переводы
	0. Выход`

func UserWin(db *sql.DB, user modules.User) {
	fmt.Println(UserChoice)
	var userCh int64
	fmt.Scan(&userCh)
	switch userCh {
	case 1:
		fmt.Println(user)
	case 2:
		AddAccount(db)
	case 3:
		var pinUs int64
		fmt.Print("Введите pin вашего счета: ")
		fmt.Scan(&pinUs)
		modules.ShowAcc(db, pinUs)
	case 4:
		TransferWin(db)

	case 0:
		os.Exit(0)
	default:
		fmt.Println("Не корректный ввод попробуйте еще раз")
	}
}

const TransfersChoice = `Выберите операцию:
	1. Перевод денег на другой счет 	2. Оплатить услугу
	3. История транзакций			4. Назад`

func TransferWin(db *sql.DB) {
	fmt.Println()
	fmt.Println(TransfersChoice)
	var transfer int64
	_, err := fmt.Scan(&transfer)
	if err != nil {
		log.Fatal("Неправильный ввод!")
	}
	switch transfer {
	case 1:
		AddTransaction(db)
	case 2:
		SrvcWin(db)
	case 3:
		var accNumber int64
		fmt.Println("Введите номер вашего аккаунта: ")
		fmt.Scan(&accNumber)
		var pincode int64
		fmt.Print("Введите PIN-код вашего счета: ")
		fmt.Scan(&pincode)
		modules.HistoryOfTransactions(db, accNumber, pincode)
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Не корректный ввод")
	}
}

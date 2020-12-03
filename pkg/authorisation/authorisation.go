package authorisation

import (
	"database/sql"
	"fmt"
)

const FirstWindow = `"BANK-T"
 	1. Авторизация
	2. Выход`

const AdminChoice = `Выберите дейсвтия:
	1. Добавить пользователя
	2. Добавить валюту
	3. Добавить автомат
	4. Открыть счет`

func Authorisation (db *sql.DB)(login, password string){
	var firstChoice int64
	fmt.Println(FirstWindow)
	fmt.Scan(&firstChoice)
	switch firstChoice {
	case 1:
		fmt.Println("Введите логин и пароль: ")
		fmt.Println("логин: ")
		fmt.Scan(&login)
		fmt.Println("пароль: ")
		fmt.Scan(&password)
		fmt.Println(AdminChoice)
		var ch int64
		fmt.Scan(&ch)
		if(ch == 1) {
			AddUser(db)
		}
		if(ch == 2) {
			AddCurrency(db)
		}
		if(ch == 3) {
			AddAtm(db)
		}
		if(ch == 4) {
			AddAccount(db)
		}
	case 2:
		fmt.Println("До свидания!!!")
	default:
		fmt.Println("Не корректный ввод попробуйте еще раз")
	}
	return
}


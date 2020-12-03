package authorisation

import (
	"bank-t/pkg/modules"
	"database/sql"
	"fmt"
)

func AddUser(db *sql.DB)  {
	fmt.Println("Введите данные: ")
	var fname, lname, gender, login, pass, role string
	role = "user"
	var age int64
	var remove bool
	remove = true
	fmt.Println("Имя: ")
	fmt.Scan(&fname)
	fmt.Println("Фамилия: ")
	fmt.Scan(&lname)
	fmt.Println("Пол: ")
	fmt.Scan(&gender)
	fmt.Println("Логин: ")
	fmt.Scan(&login)
	fmt.Println("Пароль: ")
	fmt.Scan(&pass)
	fmt.Println("Возраст: ")
	fmt.Scan(&age)
	newUser := modules.User{
		Id:        0,
		FirstName: fname,
		LastName:  lname,
		Age:       age,
		Gender:    gender,
		Login:     login,
		Password:  pass,
		Role:      role,
		Remove:    remove,
	}
	_, err := modules.AddNewUser(db, newUser)
	if err != nil {
		fmt.Println("Can't insert to ATMs table new address, err is", err)
	}
}

func AddCurrency(db *sql.DB) {
	var name string
	fmt.Println("Введите название валюты")
	fmt.Scan(&name)
	currNew := modules.CurrencyB{
		Id:   0,
		Name: name,
	}
	modules.AddCurr(db, currNew)
}

func AddAtm(db *sql.DB) {
	var addr string
	var st bool
	st = true
	fmt.Println("Введите адрес: ")
	fmt.Scan(&addr)
	newAtm := modules.Atm{
		Id:      0,
		Address: addr,
		Status:  st,
	}
	modules.AddATM(db, newAtm)

}

func AddAccount (db *sql.DB){
	var usId, numbAcc, amount, curr int64
	var remove bool
	remove = true
	fmt.Println("Открытие Счета")
	fmt.Println("Введите данные: ")
	fmt.Println("Id пользователя: ")
	fmt.Scan(&usId)
	fmt.Println("Номер счета: ")
	fmt.Scan(&numbAcc)
	fmt.Println("Счет: ")
	fmt.Scan(&amount)
	fmt.Println("Id валюты: ")
	fmt.Scan(&curr)
	newAcc := modules.Account{
		Id:        0,
		UserId:    usId,
		NumberAcc: numbAcc,
		Amount:    amount,
		Currency:  curr,
		Remove:    remove,
	}
	modules.AddNewAccount(db, newAcc)
}
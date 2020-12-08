package authorisation

import (
	"bank-t/database"
	"bank-t/pkg/modules"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func GetLoginPassUserOrAdmin(db *sql.DB, login, password, role string) {
	var UserA modules.User
	_ = db.QueryRow(database.GetLoginPass, login, password, role).Scan(
		&UserA.Login,
		&UserA.Password,
		&UserA.Role,
	)
	role = UserA.Role
	switch role {
	case "admin":
		AdminWin(db, UserA)

	case "user":
		UserWin(db, UserA)

	default:
		fmt.Println("Повторите попытку!!!")
	}
}

func AddUser(db *sql.DB) {
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
	modules.FileRecordingJson(newUser)
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
	fmt.Println("Введите адрес: ")
	var s string
	fmt.Scan(&s)
	reader := bufio.NewReader(os.Stdin)
	address, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Can't read command: %v", err)
	}
	fmt.Println(s)
	sprintf := fmt.Sprintf("%s %s", s, address) // Объединяет введенные данные в одну строку
	fmt.Printf("Был добавлен АТМ по адресу: %s %s\n", s, address)
	_, err = modules.AddATM(db, sprintf)
	if err != nil {
		fmt.Println("Vse ploho")
		return
	}
	fmt.Println("Vse Ok")
}

func AddAccount(db *sql.DB) {
	var usId, numbAcc, amount, curr, pin int64
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
	fmt.Println("Введите Pin: ")
	fmt.Scan(&pin)
	newAcc := modules.Account{
		Id:        0,
		UserId:    usId,
		NumberAcc: numbAcc,
		Amount:    amount,
		Currency:  curr,
		Pin:       pin,
		Remove:    remove,
	}
	modules.AddNewAccount(db, newAcc)
}

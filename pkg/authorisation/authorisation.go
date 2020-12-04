package authorisation

import (
	"bank-t/pkg/modules"
	"database/sql"
	"fmt"
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
			var role string
			fmt.Print("Введите ваш роль - admin | user: ")
			fmt.Scan(&role)
			UserAdmin(db, login, password, role)
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

func UserAdmin(database *sql.DB, login, password, role string) {
	var UserA modules.User
	_ = database.QueryRow(`Select *From users Where (login=($1) and password=($2))and role=($3)`, login, password, role).Scan(
		&UserA.Id,
		&UserA.FirstName,
		&UserA.LastName,
		&UserA.Age,
		&UserA.Gender,
		&UserA.Login,
		&UserA.Password,
		&UserA.Role,
		&UserA.Remove)

	if UserA.Role == "admin" {
		AdminWin(database, UserA)
	}
	if UserA.Role == "user" {
		UserWin(database, UserA)
	}
}

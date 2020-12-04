package authorisation

import (
	"bank-t/pkg/modules"
	"database/sql"
	"fmt"
	"os"
)

const UserChoice = `Выберите дейсвтия:
	1. Посмотреть мои данные
	2. Открыть счет
	3. Мои счета

	4. Перевод денег
	5. Оплатить услугу
	6. История транзакций
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
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Не корректный ввод попробуйте еще раз")
	}
}

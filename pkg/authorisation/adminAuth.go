package authorisation

import (
	"bank-t/pkg/modules"
	"database/sql"
	"fmt"
	"os"
)

const AdminChoice = `Выберите дейсвтия:
	1. Добавить пользователя
	2. Добавить валюту
	3. Добавить автомат
	4. Открыть счет
	5. Посмотреть мои данные
	6. Мои счета

	7. Перевод денег
	8. Оплатить услугу
	9. История транзакций
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

	case 0:
		os.Exit(0)
	default:
		fmt.Println("Не корректный ввод попробуйте еще раз")
	}
}

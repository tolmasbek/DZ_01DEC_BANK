package authorisation

import (
	"bank-t/pkg/modules"
	"database/sql"
	"fmt"
)

const UserChoice = `Выберите дейсвтия:
	1. Посмотреть мои данные
	2. Посмотреть счет
	3. Перевод со счета на счет
	4. Изменить пароль
	5. Выход`

func UserWin(db *sql.DB, user modules.User)  {
	fmt.Println(UserChoice)
	var userCh int64
	fmt.Scan(&userCh)
	switch userCh {
	case 1:
		fmt.Println(user)
	default:
		fmt.Println("Не корректный ввод попробуйте еще раз")
	}
}
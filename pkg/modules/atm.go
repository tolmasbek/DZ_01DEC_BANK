package modules

import (
	"bank-t/database"
	"database/sql"
	"fmt"
)

type Atm struct {
	Id      int64
	Address string
	Status  bool
}

func AddATM(dataBase *sql.DB, address string) (ok bool, err error) {
	_, err = dataBase.Exec(database.AddNewAtm, address)
	if err != nil {
		fmt.Println("Can't insert to ATMs table new address, err is", err)
		return false, err
	}
	return true, nil
}


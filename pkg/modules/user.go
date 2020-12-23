package modules

import (
	"bank-t/database"
	"database/sql"
	"fmt"
)

type User struct {
	Id        int64  `json:"id" 		xml:"id"`
	FirstName string `json:"first_name" xml:"first_name"`
	LastName  string `json:"last_name" 	xml:"last_name"`
	Age       int64  `json:"age" 		xml:"age"`
	Gender    string `json:"gender" 	xml:"gender"`
	Login     string `json:"login" 		xml:"login"`
	Password  string `json:"password" 	xml:"password"`
	Role      string `json:"role" 		xml:"role"`
	Remove    bool   `json:"remove" 	xml:"remove"`
}

func AddNewUser(db *sql.DB, user User) (ok bool, err error) {
	_, err = db.Exec(database.AddNewUser,
		user.FirstName,
		user.LastName,
		user.Age,
		user.Gender,
		user.Login,
		user.Password,
		user.Role,
		user.Remove)
	if err != nil {
		fmt.Println("Can't inserted, err is", err)
		return false, err
	}
	return true, nil
}

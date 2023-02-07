package user

import (
	"database/sql"
	"fmt"
)

func InsertData(db *sql.DB, username string, password string) bool {
	_, err := db.Exec("insert into user (username ,password) value (?,?)", username, password)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

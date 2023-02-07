package user

import (
	"database/sql"
)

func ChangeSpecificUserData(db *sql.DB, userName string, sqlWords, requestData string) bool {
	_, err := db.Exec(sqlWords, requestData, userName)
	if err != nil {
		return false
	}
	return true
}

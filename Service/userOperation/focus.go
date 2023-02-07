package userOperation

import (
	"bookshop/DataHandler/userOperation"
	"bookshop/Service/user"
	"database/sql"
)

func FocusWithToken(db *sql.DB, token string, focusedUserId string) (int, map[string]interface{}) {
	data, err := user.ParseToken(token)
	if err != nil {
		return 202, map[string]interface{}{
			"info":   "auth fail",
			"status": 1,
		}
	}
	if focusedUserId == "" {
		return 202, map[string]interface{}{
			"info":   "no focusedUserId",
			"status": 3,
		}
	}

	username := data.Issuer
	ok := userOperation.FocusWithFocusUserNameAndFocusedUserId(db, username, focusedUserId)
	if !ok {
		return 202, map[string]interface{}{
			"info":   "insert data fail",
			"status": 2,
		}
	}
	return 200, map[string]interface{}{
		"info":   "success",
		"status": 0,
	}
}

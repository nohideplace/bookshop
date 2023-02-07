package userOperation

import (
	"bookshop/Model"
	"database/sql"
	"fmt"
)

func FocusWithFocusUserNameAndFocusedUserId(db *sql.DB, username string, focusedUserId string) bool {
	var user Model.User
	userRows, err1 := db.Query("select id from user where username=?", username)
	if err1 != nil {
		fmt.Println(err1)
		return false
	}
	if userRows.Next() {
		err2 := userRows.Scan(&user.Id)
		if err2 != nil {
			fmt.Println(err2)
			return false
		}
	} else {
		return false
	}
	//获取到用户id
	userId := user.Id
	_, err2 := db.Exec("insert into user_focus(focus_user_id, focused_user_id) value(?,?)", userId, focusedUserId)
	if err2 != nil {
		return false
	}
	return true
}

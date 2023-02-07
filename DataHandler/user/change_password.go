package user

import "database/sql"

func UpdatePass(db *sql.DB, username string, oldPass string, newPass string) bool {
	data, err := SelectFromUserName(db, username)
	if err != nil {
		return false
	}
	selectedPass := data.Password
	//提供的密码校验不上
	if selectedPass != oldPass {
		return false
	}
	_, err1 := db.Exec("update user set password=? where username=?", newPass, username)
	if err1 != nil {
		return false
	}
	return true
}

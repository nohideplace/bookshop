package user

import (
	"bookshop/DataHandler/user"
	"database/sql"
)

func CheckAndChangePass(db *sql.DB, token string, oldPass string, newPass string) (int, map[string]interface{}) {
	if len(newPass) > 20 || len(oldPass) > 20 || len(newPass) < 3 || len(oldPass) < 3 {
		return 202, map[string]interface{}{
			"status": 5,
			"info":   "input data invalid",
		}
	}
	ok, _ := AuthCheck(token)
	//校验不通过
	if !ok {
		return 202, map[string]interface{}{
			"info":   "auth check fail",
			"status": 1,
		}
	}
	//两次密码一样就返回密码一样
	if oldPass == newPass {
		return 202, map[string]interface{}{
			"info":   "same pass",
			"status": 2,
		}
	}
	//从token中获取用户名
	claim, err := ParseToken(token)
	if err != nil {
		return 202, map[string]interface{}{
			"info":   "parse token fail",
			"status": 3,
		}
	}
	username := claim.Issuer
	if user.UpdatePass(db, username, oldPass, newPass) {
		return 200, map[string]interface{}{
			"info":   "success",
			"status": 0,
		}
	}
	return 202, map[string]interface{}{
		"info":   "origin pass wrong",
		"status": 4,
	}
}

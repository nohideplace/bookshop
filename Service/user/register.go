package user

import (
	"bookshop/DataHandler/user"
	"database/sql"
)

// CheckAndInsert 检查注册输入的合法性，合法则继续插入，注意不同错误的返回不同，在api处直接调用就可以了
func CheckAndInsert(db *sql.DB, username string, password string) (int, map[string]interface{}) {
	//检查输入的合法性
	if len(username) > 20 || len(password) > 20 || len(username) < 3 || len(password) < 3 {
		return 202, map[string]interface{}{
			"status": 3,
			"info":   "input data invalid",
		}
	}
	if username == "" || password == "" {
		return 202, map[string]interface{}{"status": 1, "info": "No username or password"}
	}
	//在检查输入合法后，尝试向数据库插入
	ok := user.InsertData(db, username, password)
	if !ok {
		return 202, map[string]interface{}{"status": 2, "info": "Insert fail"}
	}
	return 200, map[string]interface{}{"status": 0, "info": "Insert success"}
}

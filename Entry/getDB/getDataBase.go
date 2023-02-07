package getDB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	//以root用户，密码114514连接到127.0.0.1:3306的test数据库
	var dns = "root:chrnbfj666@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}

package user

import (
	"bookshop/Model"
	"database/sql"
	"errors"
	"fmt"
)

func SelectFromUserId(db *sql.DB, userId string) (Model.User, error) {
	var data Model.User
	rows, err1 := db.Query("select * from user where id=?", userId)
	if err1 != nil {
		fmt.Println(err1)
		return Model.User{}, err1
	}
	defer rows.Close()
	//后移取一位，并验证是否为空
	ok := rows.Next()
	if !ok {
		fmt.Println("没有查询到")
		return Model.User{}, errors.New("没有查询到")
	}
	err2 := rows.Scan(&data.Id, &data.Username, &data.Password, &data.Gender, &data.Nickname, &data.QQ, &data.Birthday, &data.Email, &data.Avatar, &data.Introduction, &data.Phone)
	if err2 != nil {

		fmt.Println(err2)
		return Model.User{}, err2
	}
	return data, nil
}

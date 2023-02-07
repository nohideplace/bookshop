package comment

import (
	"bookshop/Model"
	"database/sql"
	"fmt"
	"time"
)

// MakeComment 返回插入id和操作是否成功
func MakeComment(db *sql.DB, username string, bookId string, content string, parentId []string, ok bool) (int, bool) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	sqlStrWithParentId := "insert into comment(parent_id, book_id, publish_time, content, user_id)value(?,?,?,?,?)"
	sqlStrWithoutParentId := "insert into comment(book_id, publish_time, content, user_id)value(?,?,?,?)"
	var user Model.User
	userRows, err1 := db.Query("select id from user where username=?", username)
	if err1 != nil {
		fmt.Println(err1)
		return 0, false
	}
	if userRows.Next() {
		err2 := userRows.Scan(&user.Id)
		if err2 != nil {
			fmt.Println(err2)
			return 0, false
		}
	} else {
		return 0, false
	}
	//获取到用户id
	userId := user.Id
	//没有传入parentId
	if !ok || parentId[0] == "" {
		ret, err3 := db.Exec(sqlStrWithoutParentId, bookId, timeStr, content, userId)
		if err3 != nil {
			fmt.Println(err3)
			return 0, false
		}
		id, err := ret.LastInsertId()
		if err != nil {
			fmt.Println(err)
			return 0, false
		}
		return int(id), true
	} else {
		ret, err3 := db.Exec(sqlStrWithParentId, parentId[0], bookId, timeStr, content, userId)
		if err3 != nil {
			fmt.Println(err3)
			return 0, false
		}
		id, err := ret.LastInsertId()
		if err != nil {
			fmt.Println(err)
			return 0, false
		}
		return int(id), true
	}
}

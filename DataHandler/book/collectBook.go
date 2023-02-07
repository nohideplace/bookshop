package book

import (
	"database/sql"
	"fmt"
)

//收藏书，插入表中一条数据

func CollectBookWithRelativeInfo(db *sql.DB, bookName string, bookId int, userName string, userId int) bool {
	_, err := db.Exec("insert into collection (book_id, user_id, bookname, username) value(?,?,?,?)", bookId, userId, bookName, userName)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

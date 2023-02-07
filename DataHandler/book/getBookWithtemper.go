package book

import "database/sql"

// GetCollectionsCount GetCollections 获取某本书的收藏数
func GetCollectionsCount(db *sql.DB, bookId int) int {
	rows, err := db.Query("select * from collection where book_id=?", bookId)
	if err != nil {
		return 0
	}
	var counter int
	for rows.Next() {
		//每条数据自增+1
		counter++
	}
	return counter
}

func GetCommentCount(db *sql.DB, bookId int) int {
	rows, err := db.Query("select * from comment where book_id=?", bookId)
	if err != nil {
		return 0
	}
	var counter int
	for rows.Next() {
		//每条数据自增+1
		counter++
	}
	return counter
}

package book

import (
	"bookshop/DataHandler/book"
	"bookshop/Service/user"
	"database/sql"
)

func GetAllBook(db *sql.DB, originToken []string) (int, map[string]interface{}) {
	//没有提供token时，is_star字段为空
	books, err := book.GetAllBook(db)
	if err != nil {
		return 202, map[string]interface{}{
			"status": 1,
			"info":   "get books fail",
			"data": map[string]interface{}{
				"books": []map[string]interface{}{}},
		}
	}
	//当没有传入token时
	if len(originToken) == 0 {
		//is_star默认赋值为false
		return 200, map[string]interface{}{
			"status": 0,
			"info":   "success",
			"data": map[string]interface{}{
				"books": books},
		}
	}
	data, err1 := user.ParseToken(originToken[0])
	if err1 != nil {
		return 200, map[string]interface{}{
			"status": 2,
			"info":   "parse token fail, return default books(without is_star)",
			"data": map[string]interface{}{
				"books": books},
		}
	}
	username := data.Issuer
	collections, err2 := book.GetCollectionsWithUserName(db, username)
	if err2 != nil {
		return 200, map[string]interface{}{
			"status": 3,
			"info":   "get collections fail, return default books(without is_star)",
			"data": map[string]interface{}{
				"books": books},
		}
	}
	//collection表的记录是有一条就意味着收藏了
	//range是复制，不能修改到值
	for i := 0; i < len(collections); i++ {
		for j := 0; j < len(books); j++ {
			if collections[i].BookId == books[j].Id {
				books[j].IsStar = true
			}
		}
	}
	return 200, map[string]interface{}{
		"status": 0,
		"info":   "success",
		"data": map[string]interface{}{
			"books": books},
	}
}

package book

import (
	"bookshop/DataHandler/book"
	"bookshop/Service/user"
	"database/sql"
)

func GetBookWithBookName(db *sql.DB, originToken []string, bookName string) (int, map[string]interface{}) {
	//首先检查输入是否为空，如果为空，返回
	data, err1 := book.GetBookWithBookName(db, bookName)
	if err1 != nil {
		return 202, map[string]interface{}{
			"status": 1,
			"info":   "Get data fail",
			"data": map[string]interface{}{
				"books": map[string]interface{}{}},
		}
	}
	//没有token
	if len(originToken) == 0 {
		//is_star默认赋值为false
		return 200, map[string]interface{}{
			"status": 0,
			"info":   "success with no token",
			"data": map[string]interface{}{
				"books": data},
		}
	}
	//有token，先校验
	claim, err2 := user.ParseToken(originToken[0])
	if err2 != nil {
		return 200, map[string]interface{}{
			"status": 2,
			"info":   "parse token fail, return default books(without is_star)",
			"data": map[string]interface{}{
				"books": data},
		}
	}
	collection, err3 := book.GetCollectionsWithUserName(db, claim.Issuer)
	if err3 != nil {
		return 200, map[string]interface{}{
			"status": 2,
			"info":   "parse token fail, return default books(without is_star)",
			"data": map[string]interface{}{
				"books": data},
		}
	}
	for i := 0; i < len(collection); i++ {
		if collection[i].BookName == bookName {
			data.IsStar = true
		}
	}
	return 200, map[string]interface{}{
		"status": 0,
		"info":   "success",
		"data": map[string]interface{}{
			"books": data},
	}

}

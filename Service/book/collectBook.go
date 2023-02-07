package book

import (
	"bookshop/DataHandler/book"
	user2 "bookshop/DataHandler/user"
	"bookshop/Service/user"
	"database/sql"
)

//收藏书籍
/*
首先校验token是否有，必须有
然后根据bookId获取到书名
然后在用户书籍关系表插入一条信息
返回即可

*/

func CollectBook(db *sql.DB, originToken []string, bookId string) (int, map[string]interface{}) {
	//校验token不通过
	if len(originToken) == 0 {
		return 202, map[string]interface{}{
			"info":   "no auth data",
			"status": 1,
		}
	}
	ok, _ := user.AuthCheck(originToken[0])
	if !ok {
		return 202, map[string]interface{}{
			"info":   "auth check fail",
			"status": 2,
		}
	}
	data, err1 := user.ParseToken(originToken[0])
	if err1 != nil {
		return 202, map[string]interface{}{
			"info":   "parse token fail",
			"status": 3,
		}
	}
	username := data.Issuer
	userdata, err2 := user2.SelectFromUserName(db, username)
	if err2 != nil {
		return 202, map[string]interface{}{
			"info":   "select user info fail",
			"status": 4,
		}
	}
	userId := userdata.Id
	bookdata, err3 := book.GetBookWithBookId(db, bookId)
	if err3 != nil {
		return 202, map[string]interface{}{
			"info":   "select book info fail",
			"status": 5,
		}
	}
	ok2 := book.CollectBookWithRelativeInfo(db, bookdata.Name, bookdata.Id, username, userId)
	if !ok2 {
		return 202, map[string]interface{}{
			"info":   "Collect fail",
			"status": 6,
		}
	}
	return 200, map[string]interface{}{
		"info":   "success",
		"status": 0,
	}

}

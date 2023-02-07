package comment

import (
	"bookshop/DataHandler/comment"
	"bookshop/Service/user"
	"database/sql"
	"mime/multipart"
)

func MakeComment(db *sql.DB, tokenStr string, bookId string, formData *multipart.Form, err error) (int, map[string]interface{}) {
	//校验参数，bookId和content都是必选参数
	if err != nil {
		return 202, map[string]interface{}{
			"info":   "fail",
			"status": 1,
			"data":   "获取表单错误",
		}
	}
	if bookId == "" {
		return 202, map[string]interface{}{
			"info":   "fail",
			"status": 2,
			"data":   "bookId为空",
		}
	}
	tokenData, err3 := user.ParseToken(tokenStr)
	if err3 != nil {
		return 202, map[string]interface{}{
			"info":   "fail",
			"status": 2,
			"data":   "鉴权失败",
		}
	}
	username := tokenData.Issuer
	content, ok := formData.Value["content"]
	if !ok {
		return 202, map[string]interface{}{
			"info":   "fail",
			"status": 3,
			"data":   "表单中没有content参数",
		}
	}
	parentId, ok1 := formData.Value["parent_id"]
	//parent_id没有传入
	id, ok3 := comment.MakeComment(db, username, bookId, content[0], parentId, ok1)
	if !ok3 {
		return 202, map[string]interface{}{
			"info":   "fail",
			"status": 3,
			"data":   "插入数据失败",
		}
	}
	return 200, map[string]interface{}{
		"info":   "success",
		"status": 0,
		"data":   id,
	}
}

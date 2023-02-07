package comment

import (
	"bookshop/DataHandler/comment"
	"bookshop/Service/user"
	"database/sql"
	"mime/multipart"
)

func UpdateComment(db *sql.DB, token string, commentId string, data *multipart.Form, err error) (int, map[string]interface{}) {
	if err != nil {
		return 202, map[string]interface{}{
			"info":   "get form fail",
			"status": 5,
		}
	}
	ok1, _ := user.AuthCheck(token)
	if !ok1 {
		return 202, map[string]interface{}{
			"info":   "auth check fail",
			"status": 1,
		}
	}
	content, ok2 := data.Value["content"]
	if !ok2 {
		return 202, map[string]interface{}{
			"info":   "no content",
			"status": 2,
		}
	}
	ok3 := comment.UpdateComment(db, commentId, content[0])
	if !ok3 {
		return 202, map[string]interface{}{
			"info":   "update data fail",
			"status": 3,
		}
	}
	return 200, map[string]interface{}{
		"info":   "success",
		"status": 0,
	}
}

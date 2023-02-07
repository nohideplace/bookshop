package userOperation

import (
	"bookshop/DataHandler/userOperation"
	"bookshop/Service/user"
	"database/sql"
	"mime/multipart"
)

func PraiseWithCommentId(db *sql.DB, token string, formData *multipart.Form, err error) (int, map[string]interface{}) {
	if err != nil {
		return 202, map[string]interface{}{
			"info":   "no form data",
			"status": 3,
		}
	}
	data, err1 := user.ParseToken(token)
	commentId, ok1 := formData.Value["target_id"]
	if !ok1 || commentId[0] == "" {
		return 202, map[string]interface{}{
			"info":   "no commentId",
			"status": 4,
		}
	}
	if err1 != nil {
		return 202, map[string]interface{}{
			"info":   "auth check fail",
			"status": 1,
		}
	}
	username := data.Issuer
	ok2 := userOperation.Praise(db, username, commentId[0])
	if !ok2 {
		return 202, map[string]interface{}{
			"info":   "insert data fail",
			"status": 2,
		}
	}
	return 200, map[string]interface{}{
		"info":   "success",
		"status": 0,
	}
}

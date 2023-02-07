package comment

import (
	"bookshop/DataHandler/comment"
	"bookshop/Service/user"
	"database/sql"
)

func DelComment(db *sql.DB, token string, commentId string) (int, map[string]interface{}) {
	ok, _ := user.AuthCheck(token)
	if commentId == "" {
		return 202, map[string]interface{}{
			"info":   "no commentId",
			"status": 4,
		}
	}
	if !ok {
		return 202, map[string]interface{}{
			"info":   "auth check fail",
			"status": 1,
		}
	}
	ok1 := comment.DelCommentWithCommentId(db, commentId)
	if !ok1 {
		return 202, map[string]interface{}{
			"info":   "del data fail",
			"status": 2,
		}
	}
	return 200, map[string]interface{}{
		"info":   "success",
		"status": 0,
	}
}

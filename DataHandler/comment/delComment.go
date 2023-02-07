package comment

import (
	"database/sql"
	"fmt"
)

func DelCommentWithCommentId(db *sql.DB, commentId string) bool {
	_, err1 := db.Exec("delete from comment where id=?", commentId)
	if err1 != nil {
		fmt.Println(err1)
		return false
	}
	_, err2 := db.Exec("delete from comment_praise where comment_id=?", commentId)
	if err2 != nil {
		return false
	}
	return true
}

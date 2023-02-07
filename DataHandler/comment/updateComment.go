package comment

import "database/sql"

func UpdateComment(db *sql.DB, commentId string, content string) bool {
	_, err := db.Exec("update comment set content=? where id=?", content, commentId)
	if err != nil {
		return false
	}
	return true
}

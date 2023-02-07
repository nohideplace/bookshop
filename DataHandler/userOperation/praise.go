package userOperation

import (
	"bookshop/Model"
	"database/sql"
	"fmt"
)

func Praise(db *sql.DB, username string, commentId string) bool {
	//从comment_id获取到book_id，praised_user_id
	var comment Model.CommentPraise
	userRows, err1 := db.Query("select book_id, user_id from comment where id=?", commentId)
	if err1 != nil {
		fmt.Println(err1)
		return false
	}
	if userRows.Next() {
		err2 := userRows.Scan(&comment.BookId, &comment.PraisedUserId)
		if err2 != nil {
			fmt.Println(err2)
			return false
		}
	} else {
		return false
	}
	//获取到用户id
	bookId := comment.BookId
	praisedUserId := comment.PraisedUserId
	//从username获取到praise_user_id
	var user Model.User
	userRows1, err2 := db.Query("select id from user where username=?", username)
	if err2 != nil {
		fmt.Println(err2)
		return false
	}
	if userRows1.Next() {
		err3 := userRows1.Scan(&user.Id)
		if err3 != nil {
			fmt.Println(err3)
			return false
		}
	} else {
		return false
	}
	//获取到用户id
	praiseUserId := user.Id
	_, err := db.Exec("insert into comment_praise(book_id, comment_id, praise_user_id, praised_user_id)value(?,?,?,?)", bookId, commentId, praiseUserId, praisedUserId)
	if err != nil {
		return false
	}
	_, err3 := db.Exec("update comment set praise_count=praise_count+1 where id=?", commentId)
	if err3 != nil {
		return false
	}
	return true

}

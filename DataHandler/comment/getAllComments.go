package comment

import (
	"bookshop/Model"
	"database/sql"
	"errors"
	"fmt"
)

//先获取打底内容

func GetAllCommentsWithBookId(db *sql.DB, bookId string) ([]Model.Comment, error) {
	var comment Model.Comment
	var comments []Model.Comment
	rows, err1 := db.Query("select * from comment where book_id=?", bookId)
	if err1 != nil {
		fmt.Println(err1)
		return []Model.Comment{}, err1
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&comment.Id, &comment.ParentId, &comment.BookId, &comment.PublishTime, &comment.Content, &comment.UserId, &comment.Avatar, &comment.NickName, &comment.PraiseCount)
		if err != nil {
			fmt.Println(err)
			return []Model.Comment{}, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

//这两张表都没有用户名，需要先查一次用户表，获取到用户id

// GetPraiseWithUserName 获取当前用户在当前书籍下，对所有评论的点赞信息
func GetPraiseWithUserName(db *sql.DB, username string, bookId string) ([]Model.CommentPraise, error) {
	//首先根据用户名查询到id
	var user Model.User
	userRows, err1 := db.Query("select id from user where username=?", username)
	if err1 != nil {
		fmt.Println(err1)
		return []Model.CommentPraise{}, err1
	}
	if userRows.Next() {
		err2 := userRows.Scan(&user.Id)
		if err2 != nil {
			fmt.Println(err2)
			return nil, err2
		}
	} else {
		fmt.Println("没有查询到相关信息")
		return nil, errors.New("没有查询到相关信息")
	}
	//获取到用户id
	userId := user.Id
	var commentPraise Model.CommentPraise
	var commentPraises []Model.CommentPraise
	praiseRows, err3 := db.Query("select * from comment_praise where praise_user_id=? and book_id=?", userId, bookId)
	if err3 != nil {
		fmt.Println(err3)
		return nil, err3
	}
	for praiseRows.Next() {
		err4 := praiseRows.Scan(&commentPraise.Id, &commentPraise.BookId, &commentPraise.CommentId, &commentPraise.PraiseUserId, &commentPraise.PraisedUserId)
		if err4 != nil {
			fmt.Println(err4)
			return nil, err4
		}
		commentPraises = append(commentPraises, commentPraise)
	}
	return commentPraises, nil
}

// GetFocusWithUserName 获取当前用户的所有关注用户
func GetFocusWithUserName(db *sql.DB, username string) ([]Model.UserFocus, error) {
	var user Model.User
	userRows, err1 := db.Query("select id from user where username=?", username)
	if err1 != nil {
		fmt.Println(err1)
		return []Model.UserFocus{}, err1
	}
	if userRows.Next() {
		err2 := userRows.Scan(&user.Id)
		if err2 != nil {
			fmt.Println(err2)
			return nil, err2
		}
	} else {
		return nil, errors.New("没有查询到相关信息")
	}
	//获取到用户id
	userId := user.Id
	var focus Model.UserFocus
	var focuses []Model.UserFocus
	focusRows, err3 := db.Query("select * from user_focus where focus_user_id=?", userId)
	if err3 != nil {
		fmt.Println(err3)
		return nil, err3
	}
	for focusRows.Next() {
		err4 := focusRows.Scan(&focus.Id, &focus.FocusUserId, &focus.FocusedUserId)
		if err4 != nil {
			fmt.Println(err4)
			return nil, err4
		}
		focuses = append(focuses, focus)
	}

	return focuses, nil
}

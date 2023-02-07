package comment

import (
	"bookshop/DataHandler/comment"
	"bookshop/Service/user"
	"database/sql"
	"fmt"
)

//注意处理，登录和未登录的关系

func GetAllComments(db *sql.DB, originToken []string, bookId string) (int, map[string]interface{}) {
	comments, err := comment.GetAllCommentsWithBookId(db, bookId)
	if err != nil {
		return 202, map[string]interface{}{
			"status": 1,
			"info":   "get books fail",
			"data": map[string]interface{}{
				"books": []map[string]interface{}{}},
		}
	}
	//没有提供token时，是否点赞，是否关注字段为空
	if len(originToken) == 0 {
		return 200, map[string]interface{}{
			"status": 0,
			"info":   "未登录，没有",
			"data": map[string]interface{}{
				"books": comments},
		}
	}
	//解析token
	data, err1 := user.ParseToken(originToken[0])
	if err1 != nil {
		return 200, map[string]interface{}{
			"status": 2,
			"info":   "parse token fail, return default comments",
			"data": map[string]interface{}{
				"books": comments},
		}
	}
	username := data.Issuer
	//拿到用户名，去评论点赞表中查找相关信息，是否给该书点赞
	//获取到这个用户，在这本书籍下的所有点赞信息
	praises, err2 := comment.GetPraiseWithUserName(db, username, bookId)
	if err2 != nil {
		return 200, map[string]interface{}{
			"status": 3,
			"info":   "获取用户点赞失败，返回默认内容",
			"data": map[string]interface{}{
				"books": comments},
		}
	}
	//当前用户的点赞表，对比评论id，对得上就设置为已点赞
	for i := 0; i < len(comments); i++ {
		for j := 0; j < len(praises); j++ {
			if praises[j].CommentId == comments[i].Id {
				comments[i].IsPraised = true
			}
		}
	}
	//拿到用户名，去用户关注表中查找相关信息，是否关注该用户
	focus, err3 := comment.GetFocusWithUserName(db, username)
	if err3 != nil {
		fmt.Println(err3)
		return 200, map[string]interface{}{
			"status": 2,
			"info":   "获取用户关注失败，返回默认内容",
			"data": map[string]interface{}{
				"books": comments},
		}
	}
	//当前用户查关注表，获取到当前用户的所有关注用户，然后和发表这些评论的所有用户对比,如果是已关注用户，则设置为已关注
	for i := 0; i < len(comments); i++ {
		for j := 0; j < len(focus); j++ {
			if focus[j].FocusedUserId == comments[i].UserId {
				comments[i].IsFocus = true
			}
		}
	}
	return 200, map[string]interface{}{
		"status": 2,
		"info":   "success",
		"data": map[string]interface{}{
			"books": comments},
	}
}

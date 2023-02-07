package comment

import (
	"bookshop/DataHandler/comment"
	"bookshop/Model"
	"bookshop/Service/user"
	"database/sql"
	"fmt"
)

func GetBetterComments(db *sql.DB, originToken []string, bookId string) (int, map[string]interface{}) {
	comments, err := comment.GetAllCommentsWithBookId(db, bookId)
	if err != nil {
		fmt.Println(err)
		return 202, map[string]interface{}{
			"status": 1,
			"info":   "get books fail",
			"data": map[string]interface{}{
				"books": []map[string]interface{}{}},
		}
	}
	//相比直接获取书本列表，这里使用一个算法进行排序
	/*
		1.遍历一遍所有书本，获取到所有parentId为0的根评论
		2.构造一个复杂类型的二维数组
		3.内部的每个一维数组就是一个从根评论开始的嵌套评论之一
	*/
	var roots []Model.Comment
	for _, book := range comments {
		if book.ParentId == 0 {
			roots = append(roots, book)
		}
	}
	var betterComments [][]Model.Comment
	for _, rootComment := range roots {
		var betterComment []Model.Comment
		//每个嵌套评论的第一条就是根评论
		betterComment = append(betterComment, rootComment)
		//从根向下寻找，所有parent_id为root_id的
		for _, com := range comments {
			if com.ParentId == rootComment.Id {
				betterComment = append(betterComment, com)
			}
		}
		betterComments = append(betterComments, betterComment)
	}

	//没有提供token时，是否点赞，是否关注字段为空
	if len(originToken) == 0 {
		return 200, map[string]interface{}{
			"status": 0,
			"info":   "未登录，部分字段默认为空",
			"data": map[string]interface{}{
				"books": betterComments},
		}
	}
	//解析token
	data, err1 := user.ParseToken(originToken[0])
	if err1 != nil {
		return 200, map[string]interface{}{
			"status": 2,
			"info":   "parse token fail, return default comments",
			"data": map[string]interface{}{
				"books": betterComments},
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
				"books": betterComments},
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
		return 202, map[string]interface{}{
			"status": 2,
			"info":   "获取用户关注失败，返回默认内容",
			"data": map[string]interface{}{
				"books": betterComments},
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

	var roots1 []Model.Comment
	for _, book := range comments {
		if book.ParentId == 0 {
			roots1 = append(roots1, book)
		}
	}
	var betterComments1 [][]Model.Comment
	for _, rootComment := range roots1 {
		var betterComment1 []Model.Comment
		//每个嵌套评论的第一条就是根评论
		betterComment1 = append(betterComment1, rootComment)
		//从根向下寻找，所有parent_id为root_id的
		for _, com := range comments {
			if com.ParentId == rootComment.Id {
				betterComment1 = append(betterComment1, com)
			}
		}
		betterComments1 = append(betterComments1, betterComment1)
	}
	return 200, map[string]interface{}{
		"status": 0,
		"info":   "success",
		"data": map[string]interface{}{
			"books": betterComments1},
	}
}

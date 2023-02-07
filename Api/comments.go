package Api

import (
	comment2 "bookshop/Service/comment"
	"database/sql"
	"github.com/gin-gonic/gin"
)

/*
二.业务流程：
1.所有用户可以进入评论区，查看所有用户的发言
2.用户可以向评论区发言
(1)发言时附带评论id，则将该评论的id作为自己的父id
(2)发言时不带评论id，传入114514，作为一条新的根评论，父评论id设置为114514
3.遍历评论区，
(1)首先找到所有父id为114514的评论，从它们开始往下递归，每个这样的评论都存在一个切片内
(2)递归在数据表中查找所有父id为它的评论，每找到一个就往切片中append一个元素，最后遍历完毕后，将这个切片append到(1)中的切片内
(3)最后返回一个装满所有评论的结构体切片的切片，还有执行成功或失败的返回
*/

func Comments(db *sql.DB, router *gin.Engine) *gin.Engine {
	router.GET("/comment/:book_id", getAllComments(db), func(c *gin.Context) {
	})
	router.POST("/comment/:book_id", makeComment(db), func(c *gin.Context) {
	})
	router.DELETE("/comment/:comment_id", delComment(db), func(c *gin.Context) {
	})
	router.PUT("/comment/:comment_id", updateComment(db), func(c *gin.Context) {
	})
	router.GET("/better_comment/:book_id", betterComments(db), func(c *gin.Context) {
	})

	return router
}

// 获取某本书下的所有评论
func getAllComments(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		originToken := c.Request.Header.Values("Authorization")
		bookId := c.Param("book_id")
		code, resp := comment2.GetAllComments(db, originToken, bookId)
		c.JSON(code, resp)
	}
}
func makeComment(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		bookId := c.Param("book_id")
		formData, err := c.MultipartForm()
		code, resp := comment2.MakeComment(db, token, bookId, formData, err)
		c.JSON(code, resp)

	}
}
func delComment(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		bookId := c.Param("comment_id")
		//此两项为必传，不传直接报错
		code, resp := comment2.DelComment(db, token, bookId)
		c.JSON(code, resp)
	}
}
func updateComment(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		bookId := c.Param("comment_id")
		formData, err := c.MultipartForm()
		code, resp := comment2.UpdateComment(db, token, bookId, formData, err)
		c.JSON(code, resp)
	}
}

// 嵌套评论的展示
func betterComments(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		originToken := c.Request.Header.Values("Authorization")
		bookId := c.Param("book_id")
		code, resp := comment2.GetBetterComments(db, originToken, bookId)
		c.JSON(code, resp)
	}
}

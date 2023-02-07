package Api

import (
	"bookshop/Service/book"
	"database/sql"
	"github.com/gin-gonic/gin"
)

/*
本文件用于编写书籍相关接口
1.获取书籍列表：时间顺序
2.搜索书籍
3.收藏书籍
4.获取某个标签的书籍列表
5.获取书籍列表：按照热度排序
*/

func Books(db *sql.DB, router *gin.Engine) *gin.Engine {
	router.GET("/book/list", getAllBook(db), func(c *gin.Context) {
	})
	router.GET("/book/search", searchBook(db), func(c *gin.Context) {
	})
	router.PUT("/book/star", collectBook(db), func(c *gin.Context) {
	})
	router.GET("/book/label", getAllBookWithLabel(db), func(c *gin.Context) {
	})
	router.GET("/book/temper", getAllBookWithTemper(db), func(c *gin.Context) {
	})
	return router
}

func getAllBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		//对头进行校验，都能获取到书籍列表，但是没有token的用户获取到的信息中的收藏栏全是false
		rt := c.Request.Header.Values("Authorization")
		code, resp := book.GetAllBook(db, rt)
		c.JSON(code, resp)
	}
}
func searchBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rt := c.Request.Header.Values("Authorization")
		bookName := c.DefaultQuery("book_name", "")
		code, resp := book.GetBookWithBookName(db, rt, bookName)
		c.JSON(code, resp)
	}
}
func collectBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		originToken := c.Request.Header.Values("Authorization")
		bookId := c.DefaultPostForm("book_id", "")
		code, resp := book.CollectBook(db, originToken, bookId)
		c.JSON(code, resp)
	}
}
func getAllBookWithLabel(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		originToken := c.Request.Header.Values("Authorization")
		label := c.DefaultQuery("label", "")
		code, resp := book.GetAllBookWithLabel(db, originToken, label)
		c.JSON(code, resp)
	}
}

// 利用热度算法，返回经过热度算法排序后的书籍
func getAllBookWithTemper(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//需要获取时间，收藏数等参数，用于计算热度，计算出的热度值赋予每本书，然后每本书按照热度值排序
		//传入书籍的id，获取到书籍的出版时间，收藏数，评论总数，综合计算赋予热度值
		originToken := c.Request.Header.Values("Authorization")
		code, resp := book.GetAllBookWithTemper(db, originToken)
		c.JSON(code, resp)
	}
}

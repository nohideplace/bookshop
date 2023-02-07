package Api

import (
	"bookshop/Service/userOperation"
	"database/sql"
	"github.com/gin-gonic/gin"
)

/*
此文件用于编写用户的操作接口
1.点赞
2.获取用户收藏列表
3.关注用户
*/

func UserOperation(db *sql.DB, router *gin.Engine) *gin.Engine {
	router.PUT("/operate/praise", praise(db), func(c *gin.Context) {
	})
	router.GET("/operate/collect/list", getList(db), func(c *gin.Context) {
	})
	router.PUT("/operate/focus", focus(db), func(c *gin.Context) {
	})
	return router
}

func praise(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		formData, err := c.MultipartForm()
		code, resp := userOperation.PraiseWithCommentId(db, token, formData, err)
		c.JSON(code, resp)
	}
}
func getList(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//传入token
		token := c.GetHeader("Authorization")
		code, resp := userOperation.GetCollectionsList(db, token)
		c.JSON(code, resp)
	}
}
func focus(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		focusedUserId := c.DefaultPostForm("user_id", "")
		code, resp := userOperation.FocusWithToken(db, token, focusedUserId)
		c.JSON(code, resp)
	}
}

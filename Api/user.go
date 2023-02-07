package Api

import (
	"bookshop/Service/user"
	"database/sql"
	"github.com/gin-gonic/gin"
)

/*
一.此文件编写用户逻辑
1.注册
2.登录：获取token
3.刷新token
4.修改用户密码
5.获取用户信息
6.修改用户信息
*/

func Users(db *sql.DB, router *gin.Engine) *gin.Engine {
	router.POST("/register", register(db), func(c *gin.Context) {
	})
	router.GET("/user/token", login(db), func(c *gin.Context) {
	})
	router.GET("/user/token/refresh", refreshToken(db), func(c *gin.Context) {
	})
	router.PUT("/user/password", changePasswd(db), func(c *gin.Context) {
	})
	router.GET("/user/info/:user_id", getUserInfo(db), func(c *gin.Context) {
	})
	router.PUT("/user/info", changeUserInfo(db), func(c *gin.Context) {
	})

	return router
}

// 用户注册：提供账号，密码
func register(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//先获取到用户名和密码
		username := c.DefaultQuery("username", "")
		password := c.DefaultQuery("password", "")
		//注意需要校验用户名和密码的长度
		code, resp := user.CheckAndInsert(db, username, password)
		c.JSON(code, resp)
	}
}
func login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取用户名和密码，然后到逻辑层去校验,校验通过，就生成token返回
		//每一个接口都要进行token校验
		username := c.DefaultQuery("username", "")
		password := c.DefaultQuery("password", "")
		//需要校验用户名和密码的长度
		code, _, _, resp := user.CheckAndReturnToken(db, username, password)
		//后面记得改成从header中获取token
		c.JSON(code, resp)
	}
}
func refreshToken(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//刷新token的原理是，先检验access_token是否合法，再检验是否过期，如果没有过期，则直接通过请求
		//如果access_token已经过期了，检查refresh_token是否合法，再检查是否过期，如果都满足，则重新生成一个access_token返回
		//如果refresh_token也不合法或者失效了，就需要用户重新登录了，并清除cookie
		//由于本人是单人后端，所以直接从cookie存放token，本来正常的前后端交互是需要把token返回就可以了，由前端将它设置到用户的cookie里，请求api时将cookie的token读出放到请求头
		//记得这个原则就好

		//获取refresh_token,然后检查refresh_token来刷新token

		token := c.GetHeader("Authorization")
		refresh_token := c.DefaultQuery("refresh_token", "")

		code, _, _, resp := user.GetTokenWithRefreshToken(token, refresh_token)
		c.JSON(code, resp)

	}
}
func changePasswd(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		oldPass := c.Query("old_password")
		newPass := c.Query("new_password")
		//密码需要校验长度
		token := c.GetHeader("Authorization")
		code, resp := user.CheckAndChangePass(db, token, oldPass, newPass)
		c.JSON(code, resp)
	}
}
func getUserInfo(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		//user_id需要校验长度
		code, resp := user.GetUserInfoWithUserId(db, user_id)
		c.JSON(code, resp)
	}
}
func changeUserInfo(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		requestData, err := c.MultipartForm()
		//表单中的每个数据都需要校验长度
		code, resp := user.ChangeUserInfo(db, accessToken, requestData, err)
		c.JSON(code, resp)
	}
}

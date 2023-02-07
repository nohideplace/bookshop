package user

import (
	"bookshop/DataHandler/user"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// CheckAndReturnToken 先查数据库，如果账号密码对的上，就生成token返回
func CheckAndReturnToken(db *sql.DB, username string, password string) (int, string, string, map[string]interface{}) {
	if len(username) > 20 || len(password) > 20 || len(username) < 3 || len(password) < 3 {
		return 202, "", "", map[string]interface{}{
			"status": 4,
			"info":   "input data invalid",
			"data":   map[string]interface{}{"refresh_token": "0", "token": "0"}}
	}
	data, err := user.SelectFromUserName(db, username)
	if err != nil {
		return 202, "", "", map[string]interface{}{
			"status": 1,
			"info":   "Select data fail",
			"data":   map[string]interface{}{"refresh_token": "0", "token": "0"}}
	}
	if username != data.Username || password != data.Password {
		return 202, "", "", map[string]interface{}{
			"status": 2,
			"info":   "Wrong username or password",
			"data":   map[string]interface{}{"refresh_token": "0", "token": "0"}}
	}
	//生成access_token和refresh_token
	token, err1 := generateToken(data.Username, 1)
	refreshToken, err2 := generateToken(data.Username, 3)
	if err1 != nil || err2 != nil {
		return 202, "", "", map[string]interface{}{
			"status": 3,
			"info":   "Generate token fail",
			"data":   map[string]interface{}{"refresh_token": "0", "token": "0"}}
	}
	return 200, token, refreshToken, map[string]interface{}{
		"status": 0,
		"info":   "success",
		"data":   map[string]interface{}{"refresh_token": refreshToken, "token": token}}
}

func generateToken(username string, duration time.Duration) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * duration).Unix(), // reqiured
		Issuer:    username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Claims = claims
	tokenString, err := token.SignedString([]byte("key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

var key = "key"

func AuthCheck(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//首先检查签名方法是否正确
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("wrong method")
		}
		//导入签名，并
		return []byte(key), nil
	})

	va, ok := err.(*jwt.ValidationError)
	//当err非空，说明令牌已经不能用了，有过期错误和令牌未激活错误，可以将其识别并返回
	if err != nil {
		if ok {
			fmt.Println(err)
			return false, err
		} else if va.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			fmt.Println("令牌已过期或未激活")
			return false, err
		} else {
			return false, err
		}
	}
	if !token.Valid {
		return false, errors.New("token不合法")
	}
	return true, nil
}

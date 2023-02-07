package user

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func GetTokenWithRefreshToken(accessToken string, refresh_token string) (int, string, string, map[string]interface{}) {
	_, err := AuthCheck(accessToken)
	//先检查通行证是否过期，过期就要重新登录获取了
	if err != nil {
		return 202, "", "", map[string]interface{}{
			"status": 1,
			"info":   "auth token fail",
			"data":   map[string]interface{}{"refresh_token": "0", "token": "0"},
		}
	}

	_, err1 := AuthCheck(refresh_token)
	//检查refreshToken是否过期，过期就重新生成一个refreshToken
	if err1 != nil {
		origin, _ := ParseToken(refresh_token)
		token, err2 := generateToken(origin.Issuer, 1)
		new_refresh_token, err3 := generateToken(origin.Issuer, 3)
		if err3 != nil || err2 != nil {
			return 202, "", "", map[string]interface{}{
				"status": 2,
				"info":   "generate token fail",
				"data":   map[string]interface{}{"refresh_token": "0", "token": "0"},
			}
		}
		return 200, new_refresh_token, token, map[string]interface{}{
			"status": 0,
			"info":   "success, return new refresh token",
			"data":   map[string]interface{}{"refresh_token": new_refresh_token, "token": token},
		}
	}
	origin, _ := ParseToken(refresh_token)
	token, err2 := generateToken(origin.Issuer, 1)
	//只生成token，不刷新refresh_token
	if err2 != nil {
		return 202, "", "", map[string]interface{}{
			"status": 3,
			"info":   "generate token fail",
			"data":   map[string]interface{}{"refresh_token": "0", "token": "0"},
		}
	}
	//没过期的情况
	return 200, refresh_token, token, map[string]interface{}{
		"status": 0,
		"info":   "success",
		"data":   map[string]interface{}{"refresh_token": refresh_token, "token": token},
	}
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	fmt.Println(tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		//首先检查签名方法是否正确
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("wrong method")
		}
		//导入签名，并
		return []byte(key), nil
	})
	origin, _ := token.Claims.(*jwt.StandardClaims)
	fmt.Println(origin)
	return origin, err
}

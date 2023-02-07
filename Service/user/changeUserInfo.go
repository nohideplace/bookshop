package user

import (
	"bookshop/DataHandler/user"
	"database/sql"
	"fmt"
	"mime/multipart"
)

var mapData = map[string]string{
	"nickname":     "update user set nickname=? where username=?",
	"avatar":       "update user set avatar=? where username=?",
	"introduction": "update user set introduction=? where username=?",
	"telephone":    "update user set telephone=? where username=?",
	"qq":           "update user set qq=? where username=?",
	"gender":       "update user set gender=? where username=?",
	"email":        "update user set email=? where username=?",
	"birthday":     "update user set birthday=? where username=?",
}

func ChangeUserInfo(db *sql.DB, token string, requestData *multipart.Form, err error) (int, map[string]interface{}) {
	//首先检查err，是否获取表单失败

	if err != nil {
		return 202, map[string]interface{}{
			"info":   "get form fail",
			"status": 1,
		}
	}
	//然后检查token是否过期
	_, err1 := AuthCheck(token)
	if err1 != nil {
		return 202, map[string]interface{}{
			"info":   "auth check fail",
			"status": 2,
		}
	}
	tokenData, _ := ParseToken(token)
	fmt.Println(tokenData.Issuer)
	//从token中提取用户id
	userName := tokenData.Issuer
	//然后利用command, ok表达式获取表单内容
	nickname, ok1 := requestData.Value["nickname"]
	avatar, ok2 := requestData.Value["avatar"]
	introduction, ok3 := requestData.Value["introduction"]
	telephone, ok4 := requestData.Value["telephone"]
	qq, ok5 := requestData.Value["qq"]
	gender, ok6 := requestData.Value["gender"]
	email, ok7 := requestData.Value["email"]
	birthday, ok8 := requestData.Value["birthday"]

	defaultResp := map[string]interface{}{
		"info":   "某项数据修改失败",
		"status": 1,
	}
	wrongInputResp := map[string]interface{}{
		"info":   "数据输入有误",
		"status": 2,
	}
	counter := 0
	switch {
	case ok1:
		if len(nickname[0]) > 15 {
			return 202, wrongInputResp
		}
		ok := user.ChangeSpecificUserData(db, userName, mapData["nickname"], nickname[0])
		if !ok {
			return 202, defaultResp
		}
		counter++
	case ok2:
		if len(avatar[0]) > 200 {
			return 202, wrongInputResp
		}
		ok := user.ChangeSpecificUserData(db, userName, mapData["avatar"], avatar[0])
		if !ok {
			return 202, defaultResp
		}
		counter++
	case ok3:
		if len(introduction[0]) > 100 {
			return 202, wrongInputResp
		}
		ok := user.ChangeSpecificUserData(db, userName, mapData["introduction"], introduction[0])
		if !ok {
			return 202, defaultResp
		}
		counter++
	case ok4:
		if len(telephone[0]) > 15 {
			return 202, wrongInputResp
		}
		ok := user.ChangeSpecificUserData(db, userName, mapData["telephone"], telephone[0])
		if !ok {
			return 202, defaultResp
		}
		counter++
	case ok5:
		if len(qq[0]) > 12 {
			return 202, wrongInputResp
		}
		ok := user.ChangeSpecificUserData(db, userName, mapData["qq"], qq[0])
		fmt.Println(qq[0])
		if !ok {
			return 202, defaultResp
		}
		counter++
	case ok6:
		if len(gender[0]) > 5 {
			return 202, wrongInputResp
		}
		ok := user.ChangeSpecificUserData(db, userName, mapData["gender"], gender[0])
		if !ok {
			return 202, defaultResp
		}
		counter++
	case ok7:
		if len(email[0]) > 30 {
			return 202, wrongInputResp
		}
		ok := user.ChangeSpecificUserData(db, userName, mapData["email"], email[0])
		if !ok {
			return 202, defaultResp
		}
		counter++
	case ok8:
		if len(birthday[0]) > 10 {
			return 202, wrongInputResp
		}
		ok := user.ChangeSpecificUserData(db, userName, mapData["birthday"], birthday[0])
		if !ok {
			return 202, defaultResp
		}
		counter++
	}
	if counter == 0 {
		return 202, map[string]interface{}{
			"info":   "no data inserted",
			"status": 2,
		}
	}

	//然后利用一个switch语句逐个检测并修改数据库内容
	return 200, map[string]interface{}{
		"info":   "success",
		"status": 0,
	}
}

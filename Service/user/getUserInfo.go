package user

import (
	"bookshop/DataHandler/user"
	"database/sql"
)

func GetUserInfoWithUserId(db *sql.DB, userId string) (int, map[string]interface{}) {
	//无需校验token

	data, err := user.SelectFromUserId(db, userId)
	if err != nil {
		return 202, map[string]interface{}{
			"status": 1,
			"info":   "fail",
			"data": map[string]interface{}{
				"user": map[string]interface{}{},
			},
		}
	}
	return 200, map[string]interface{}{
		"status": 0,
		"info":   "success",
		"data": map[string]interface{}{
			"user": map[string]interface{}{
				"id":           data.Id,
				"avatar":       data.Avatar,
				"nickname":     data.Nickname,
				"introduction": data.Introduction,
				"phone":        data.Phone,
				"qq":           data.QQ,
				"gender":       data.Gender,
				"email":        data.Email,
				"birthday":     data.Birthday,
			},
		},
	}
}

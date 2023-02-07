package userOperation

import (
	"bookshop/DataHandler/book"
	"bookshop/Model"
	"bookshop/Service/user"
	"database/sql"
)

func GetCollectionsList(db *sql.DB, token string) (int, map[string]interface{}) {
	data, err := user.ParseToken(token)
	if err != nil {
		return 202, map[string]interface{}{
			"status": 1,
			"info":   "auth fail",
			"data":   map[string]interface{}{},
		}
	}
	username := data.Issuer
	books, err1 := book.GetAllBook(db)
	if err1 != nil {
		return 202, map[string]interface{}{
			"status": 2,
			"info":   "get books fail",
			"data":   map[string]interface{}{},
		}
	}
	collections, err2 := book.GetCollectionsWithUserName(db, username)
	if err2 != nil {
		return 202, map[string]interface{}{
			"status": 2,
			"info":   "get collections fail",
			"data":   map[string]interface{}{},
		}
	}
	var collectedBooks []Model.Book
	for i := 0; i < len(collections); i++ {
		for j := 0; j < len(books); j++ {
			if collections[i].BookId == books[j].Id {
				collectedBooks = append(collectedBooks, books[j])
				//每一轮只会有一本对得上
				break
			}

		}
	}
	return 200, map[string]interface{}{
		"status": 0,
		"info":   "success",
		"data":   collectedBooks,
	}

}

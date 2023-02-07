package book

import (
	"bookshop/Model"
	"database/sql"
	"fmt"
)

func GetAllBook(db *sql.DB) ([]Model.Book, error) {
	var book Model.Book
	var books []Model.Book
	rows, err1 := db.Query("select * from book")
	if err1 != nil {

		fmt.Println(err1)
		return []Model.Book{}, err1
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Score, &book.Cover, &book.PublishTime, &book.Link, &book.Label)
		if err != nil {
			return []Model.Book{}, err
		}
		books = append(books, book)
	}
	return books, nil
}

//查找用户与书籍的收藏关系表，在这个表中查询用户某个用户名的所有书籍，检查这个用户是否收藏了这本书，收藏了就有记录，没有收藏就无记录，没有记录的在Service层返回false就可以

func GetCollectionsWithUserName(db *sql.DB, username string) ([]Model.Collection, error) {
	var collection Model.Collection
	var collections []Model.Collection
	rows, err1 := db.Query("select * from collection where username=?", username)
	if err1 != nil {
		fmt.Println(err1)
		return []Model.Collection{}, err1
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&collection.Id, &collection.BookId, &collection.UserId, &collection.UserName, &collection.BookName)
		if err != nil {
			return []Model.Collection{}, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

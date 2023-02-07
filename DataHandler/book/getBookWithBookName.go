package book

import (
	"bookshop/Model"
	"database/sql"
	"errors"
	"fmt"
)

func GetBookWithBookName(db *sql.DB, bookName string) (Model.Book, error) {
	var book Model.Book
	rows, err1 := db.Query("select * from book where name=?", bookName)
	if err1 != nil {
		fmt.Println(err1)
		return Model.Book{}, err1
	}
	defer rows.Close()
	ok := rows.Next()
	if !ok {
		fmt.Println("查询为空")
		return Model.Book{}, errors.New("查询为空")
	}
	err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Score, &book.Cover, &book.PublishTime, &book.Link, &book.Label)
	if err != nil {
		return Model.Book{}, err
	}
	return book, nil
}
func GetBookWithBookId(db *sql.DB, bookId string) (Model.Book, error) {
	var book Model.Book
	rows, err1 := db.Query("select * from book where book_id=?", bookId)
	if err1 != nil {
		fmt.Println(err1)
		return Model.Book{}, err1
	}
	defer rows.Close()
	ok := rows.Next()
	if !ok {
		fmt.Println("查询为空")
		return Model.Book{}, errors.New("查询为空")
	}
	err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Score, &book.Cover, &book.PublishTime, &book.Link, &book.Label)
	if err != nil {
		return Model.Book{}, err
	}
	return book, nil
}

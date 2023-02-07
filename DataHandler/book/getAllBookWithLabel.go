package book

import (
	"bookshop/Model"
	"database/sql"
	"fmt"
)

func GetAllBookWithLabel(db *sql.DB, label string) ([]Model.Book, error) {
	var book Model.Book
	var books []Model.Book
	rows, err1 := db.Query("select * from book where label=?", label)
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

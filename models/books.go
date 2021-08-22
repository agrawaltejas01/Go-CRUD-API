package models

import (
	"encoding/json"
	"io"
)

type Book struct {
	ID     int    `json:"id" gorm:"unique_index"`
	Title  string `json:"title" gorm:"type:varchar(50)"`
	Author string `json:"author" gorm:"type:varchar(50)"`
}

var books []Book

func AddBooks(b io.ReadCloser) (int, error) {

	var newBook Book
	_ = json.NewDecoder(b).Decode(&newBook)
	books = append(books, newBook)
	return insertBook(&newBook)

}

func GetBooks() ([]Book, error) {
	return getBooks()
}

func GetBook(id int) (Book, error) {
	return getBook(id)
}

func UpdateBook(b io.ReadCloser) error {
	var bookToUpdate Book
	_ = json.NewDecoder(b).Decode(&bookToUpdate)

	return updateBook(&bookToUpdate)
}

func DeleteBook(b io.ReadCloser) error {
	var bookToDelete Book
	_ = json.NewDecoder(b).Decode(&bookToDelete)

	return deleteBook(&bookToDelete)
}

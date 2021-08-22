package service

import (
	"data"
	"encoding/json"
	"io"
)

type Book = data.Book

func AddBooks(b io.ReadCloser) (int, error) {

	var newBook Book
	_ = json.NewDecoder(b).Decode(&newBook)
	return data.InsertBook(&newBook)

}

func GetBooks() ([]Book, error) {
	return data.GetBooks()
}

func GetBook(id int) (Book, error) {
	return data.GetBook(id)
}

func UpdateBook(b io.ReadCloser) error {
	var bookToUpdate Book
	_ = json.NewDecoder(b).Decode(&bookToUpdate)

	return data.UpdateBook(&bookToUpdate)
}

func DeleteBook(b io.ReadCloser) error {
	var bookToDelete Book
	_ = json.NewDecoder(b).Decode(&bookToDelete)

	return data.DeleteBook(&bookToDelete)
}

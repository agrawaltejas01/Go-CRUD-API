package models

import (
	"encoding/json"
	"errors"
	"io"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func AddBooks(b io.ReadCloser) []Book {

	var newBook Book
	_ = json.NewDecoder(b).Decode(&newBook)
	books = append(books, newBook)
	return books
}

func GetBooks() []Book {
	return books
}

func UpdateBook(b io.ReadCloser) ([]Book, error) {
	var bookToUpdate Book
	_ = json.NewDecoder(b).Decode(&bookToUpdate);

	updated := false
	for i := range books {
		ele := &books[i]
		if ele.ID == bookToUpdate.ID {
			ele.Author = bookToUpdate.Author
			ele.Title = bookToUpdate.Title
			updated = true
			break
		}
	}

	if !updated {
		return nil, errors.New("no book with matching ID")
	}

	return books, nil
}

func DeleteBook(b io.ReadCloser) ([]Book, error) {
	var bookToDelete Book
	_ = json.NewDecoder(b).Decode(&bookToDelete);

	deleted := false
	for i := range books {
		ele := books[i]
		if ele.ID == bookToDelete.ID {
			books = append(books[:i], books[i+1:]...)
			deleted = true
			break			
		}
	}

	if !deleted {
		return nil, errors.New("no book with matching ID")
	}

	return books, nil
}
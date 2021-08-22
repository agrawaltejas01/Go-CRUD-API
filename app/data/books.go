package data

import (
	"models"
)

type Book = models.Book

func InsertBook(book *Book) (int, error) {
	var db = models.GetDBInstance()
	result := db.Create(&book)

	if result.Error != nil {
		return 0, result.Error
	} else {
		return int(result.RowsAffected), nil
	}
}

func GetBooks() ([]Book, error) {

	var db = models.GetDBInstance()

	var books []Book
	result := db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return books, nil
	}
}

func GetBook(id int) (Book, error) {
	var db = models.GetDBInstance()

	var book Book
	result := db.First(&book, id)

	if result.Error != nil {
		return book, result.Error
	} else {
		return book, nil
	}
}

func UpdateBook(book *Book) error {
	var db = models.GetDBInstance()

	result := db.Save(&book)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteBook(book *Book) error {
	var db = models.GetDBInstance()

	result := db.Delete(&book)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

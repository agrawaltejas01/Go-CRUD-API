package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func ConnectDB() {

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	db, err = gorm.Open(dialect, dbUri)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to DB")
	}
}

func CloseDB() {
	db.Close()
}

func MigrateModels() {
	db.AutoMigrate(&Book{})
}

func insertBook(book *Book) (int, error) {
	result := db.Create(&book)

	if result.Error != nil {
		return 0, result.Error
	} else {
		return int(result.RowsAffected), nil
	}
}

func getBooks() ([]Book, error) {
	var books []Book
	result := db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return books, nil
	}
}

func getBook(id int) (Book, error) {
	var book Book
	result := db.First(&book, id)

	if result.Error != nil {
		return book, result.Error
	} else {
		return book, nil
	}
}

func updateBook(book *Book) error {
	result := db.Save(&book)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func deleteBook(book *Book) error {
	result := db.Delete(&book)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

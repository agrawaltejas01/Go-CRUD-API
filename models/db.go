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

func GetDBInstance() *gorm.DB {
	return db
}

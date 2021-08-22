package main

import (
	"controller"
	"log"
	"models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	models.ConnectDB()
	models.MigrateModels()
	defer models.CloseDB()

	router.HandleFunc("/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/book", controller.AddBook).Methods("POST")
	router.HandleFunc("/book", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book", controller.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8764", router))
}

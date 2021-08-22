package main

import (
	"log"
	"models"
	"net/http"
	"routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	models.ConnectDB()
	models.MigrateModels()
	defer models.CloseDB()

	r.HandleFunc("/books", routes.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", routes.GetBook).Methods("GET")
	r.HandleFunc("/book", routes.AddBook).Methods("POST")
	r.HandleFunc("/book", routes.UpdateBook).Methods("PUT")
	r.HandleFunc("/book", routes.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8764", r))
}

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
	r.HandleFunc("/books", routes.AddBook).Methods("POST")
	r.HandleFunc("/books", routes.UpdateBook).Methods("PUT")
	r.HandleFunc("/books", routes.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8764", r))
}

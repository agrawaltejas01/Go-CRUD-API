package main

import (
	"github.com/gorilla/mux"
	"log"
	"models"
	"net/http"
	"route"
)

func main() {
	router := mux.NewRouter()
	router.Use(route.CommonMiddleware)

	models.ConnectDB()
	models.MigrateModels()
	defer models.CloseDB()

	route.BookRoutes(router)

	log.Fatal(http.ListenAndServe(":8764", router))
}

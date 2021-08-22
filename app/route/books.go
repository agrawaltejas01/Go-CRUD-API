package route

import (
	"controller"
	"github.com/gorilla/mux"
)

func BookRoutes(router *mux.Router) {
	router.HandleFunc("/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/book", controller.AddBook).Methods("POST")
	router.HandleFunc("/book", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book", controller.DeleteBook).Methods("DELETE")

}

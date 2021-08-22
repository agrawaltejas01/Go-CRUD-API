package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"service"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	result, err := service.GetBooks()

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}
	json.NewEncoder(w).Encode(result)
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result, err := service.GetBook(id)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}
	fmt.Println("Data Found")
	json.NewEncoder(w).Encode(result)
}

func AddBook(w http.ResponseWriter, r *http.Request) {

	result, err := service.AddBooks(r.Body)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}
	json.NewEncoder(w).Encode(result)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	err := service.UpdateBook(r.Body)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}

	fmt.Println("Data Updated")
	json.NewEncoder(w).Encode(true)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	err := service.DeleteBook(r.Body)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}

	fmt.Println("Data Deleted")
	json.NewEncoder(w).Encode(true)

}

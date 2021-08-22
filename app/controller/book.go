package controller

import (
	"encoding/json"
	"net/http"
	"service"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := service.GetBooks()

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	json.NewEncoder(w).Encode(result)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result, err := service.GetBook(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	json.NewEncoder(w).Encode(result)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := service.AddBooks(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	json.NewEncoder(w).Encode(result)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := service.UpdateBook(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	json.NewEncoder(w).Encode(true)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := service.DeleteBook(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	json.NewEncoder(w).Encode(true)

}

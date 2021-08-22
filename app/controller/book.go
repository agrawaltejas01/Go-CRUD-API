package controller

import (
	"encoding/json"
	"models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := models.GetBooks()

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	json.NewEncoder(w).Encode(result)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result, err := models.GetBook(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	json.NewEncoder(w).Encode(result)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := models.AddBooks(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	json.NewEncoder(w).Encode(result)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := models.UpdateBook(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	json.NewEncoder(w).Encode(true)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := models.DeleteBook(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	json.NewEncoder(w).Encode(true)

}

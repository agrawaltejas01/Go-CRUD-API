package routes

import (
	"encoding/json"
	"models"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-Type", "application/json")
	json.NewEncoder(w).Encode(models.GetBooks())
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.AddBooks(r.Body))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books, err := models.UpdateBook(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	json.NewEncoder(w).Encode(books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books, err := models.DeleteBook(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	json.NewEncoder(w).Encode(books)

}

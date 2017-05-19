package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-inventory/models"
)

func writeJSON(w http.ResponseWriter, response models.Response) (err error) {
	res, err := json.Marshal(response)

	w.Header().Set("content-type", "application/json")
	if err != nil {
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		errRes, _ := json.Marshal(data)
		w.Write(errRes)
		return err
	}

	w.Write(res)
	return
}

// MainHandler accepts requests pointed to "/"
func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main Handled")
	data := models.Response{Err: 0, Message: "ok", Data: nil}

	writeJSON(w, data)
}

// GetCategoriesHandler returns a list of categories
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /categories")
	arr := []models.ItemCategory{
		models.ItemCategory{ID: 1, Name: "Plywood", ParentID: 0},
	}
	data := models.Response{Err: 0, Message: "", Data: arr}
	writeJSON(w, data)
}

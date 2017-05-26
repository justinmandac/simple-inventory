package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-inventory/models"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCategoriesHandler returns a list of categories
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /categories. GET params were: ", r.URL.Query())
	arr, err := categoryDao.GetCategories()

	if err != nil {
		data := models.Response{Err: 1, Message: []string{err.Error()}, Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: []string{}, Data: arr}
	writeJSON(w, data)
}

// CreateCategoryHandler creates a new category.
func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /categories.")
	var category models.ItemCategory
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&category)
	err := categoryDao.CreateCategory(category)

	if err != nil {
		data := models.Response{Err: 1, Message: []string{err.Error()}, Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: []string{}, Data: nil}
	writeJSON(w, data)
}

// DeleteCategoryHandler - deletes category
func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE /categories")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64) // the category id
	err := categoryDao.DeleteCategory(int(id))

	if err != nil {
		data := models.Response{Err: 1, Message: []string{err.Error()}, Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: []string{}, Data: nil}
	writeJSON(w, data)
}

// UpdateCategoryHandler - updates a category
func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT /categories")
	var category models.ItemCategory
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&category)
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64) // the category id

	fmt.Println("Update category id:", id)
	err := categoryDao.UpdateCategory(int(id), category)

	if err != nil {
		data := models.Response{Err: 1, Message: []string{err.Error()}, Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: []string{}, Data: nil}
	writeJSON(w, data)
}

// GetCategoryHandler - gets a single categotry
func GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	category, err := categoryDao.GetCategory(int(id))

	if err != nil {
		data := models.Response{Err: 1, Message: []string{err.Error()}, Data: nil}
		writeJSON(w, data)
		return
	}
	data := models.Response{Err: 0, Message: []string{}, Data: category}
	writeJSON(w, data)
}

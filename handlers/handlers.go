package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-inventory/models"

	"strconv"

	_ "github.com/go-sql-driver/mysql" // Import mysql driver
	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root@/inventory_v1")
	if err != nil {
		panic(err.Error())
	}
}

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

func getCategories() (categories []models.ItemCategory, err error) {
	rows, err := db.Query("SELECT * FROM categories;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var read models.ItemCategory
		err = rows.Scan(&read.ID, &read.Name, &read.ParentID)

		if err != nil {
			return nil, err
		}

		categories = append(categories, read)
	}

	return
}

// GetCategoriesHandler returns a list of categories
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /categories. GET params were: ", r.URL.Query())
	arr, err := getCategories()

	if err != nil {
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: "", Data: arr}
	writeJSON(w, data)
}

func createCategory(category models.ItemCategory) error {
	query := "INSERT INTO `categories`(`name`, `parentID`) VALUES (?, ?)"
	_, err := db.Exec(query, category.Name, category.ParentID)

	if err != nil {
		return err
	}

	return nil
}

// CreateCategoryHandler creates a new category.
func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /categories.")
	var category models.ItemCategory
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&category)
	err := createCategory(category)

	if err != nil {
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: "ok", Data: nil}
	writeJSON(w, data)
}

func deleteCategory(id int) error {
	query := "DELETE FROM `categories` WHERE id=?"
	_, err := db.Query(query, id)

	if err != nil {
		return err
	}

	return nil
}

// DeleteCategoryHandler - deletes category
func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE /categories")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64) // the category id
	err := deleteCategory(int(id))

	if err != nil {
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: "ok", Data: nil}
	writeJSON(w, data)
}

func updateCategory(id int, category models.ItemCategory) error {
	// TODO: Throw error if category.ID != id
	query := "UPDATE `categories`SET `parentID`=? WHERE `id`=?"
	_, err := db.Query(query, category.ParentID, id)

	if err != nil {
		return err
	}

	return nil
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
	err := updateCategory(int(id), category)

	if err != nil {
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: "ok", Data: nil}
	writeJSON(w, data)
}

func getCategory(id int) (cat models.ItemCategory, err error) {
	query := "SELECT * FROM `categories` WHERE id=?"
	rows, err := db.Query(query, id)
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&cat.ID, &cat.Name, &cat.ParentID)
	}

	return cat, nil
}

// GetCategoryHandler - gets a single categotry
func GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	category, err := getCategory(int(id))

	if err != nil {
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		writeJSON(w, data)
		return
	}
	data := models.Response{Err: 0, Message: "ok", Data: category}
	writeJSON(w, data)
}

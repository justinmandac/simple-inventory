package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-inventory/dao"
	"simple-inventory/models"

	"strconv"

	_ "github.com/go-sql-driver/mysql" // Import mysql driver
	"github.com/gorilla/mux"
)

var db *sql.DB
var categoryDao dao.CategoryDao
var itemDao dao.ItemDao

func init() {
	var err error
	db, err = sql.Open("mysql", "root@/inventory_v1")
	if err != nil {
		panic(err.Error())
	}
	categoryDao = dao.CategoryDao{Db: db}
	itemDao = dao.ItemDao{Db: db}
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

// GetCategoriesHandler returns a list of categories
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /categories. GET params were: ", r.URL.Query())
	arr, err := categoryDao.GetCategories()

	if err != nil {
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: "", Data: arr}
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
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: "ok", Data: nil}
	writeJSON(w, data)
}

// DeleteCategoryHandler - deletes category
func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE /categories")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64) // the category id
	err := categoryDao.DeleteCategory(int(id))

	if err != nil {
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: "ok", Data: nil}
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
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: "ok", Data: nil}
	writeJSON(w, data)
}

// GetCategoryHandler - gets a single categotry
func GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	category, err := categoryDao.GetCategory(int(id))

	if err != nil {
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		writeJSON(w, data)
		return
	}
	data := models.Response{Err: 0, Message: "ok", Data: category}
	writeJSON(w, data)
}

// GetItemsHandler retrieves the list of all items
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&item)

	err := itemDao.CreateItem(item)

	if err != nil {
		data := models.Response{Err: 1, Message: err.Error(), Data: nil}
		writeJSON(w, data)
		return
	}
	data := models.Response{Err: 0, Message: "ok", Data: nil}
	writeJSON(w, data)
}

// GetItemHandler handler for retrieving a single item
func GetItemHandler(w http.ResponseWriter, r *http.Request) {}

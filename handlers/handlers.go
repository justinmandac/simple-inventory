package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-inventory/models"

	_ "github.com/go-sql-driver/mysql" // Import mysql driver
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

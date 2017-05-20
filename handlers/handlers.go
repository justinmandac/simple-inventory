package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-inventory/models"

	"log"

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

// GetCategoriesHandler returns a list of categories
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /categories. GET params were: ", r.URL.Query())
	var arr []models.ItemCategory

	rows, err := db.Query("SELECT * FROM categories;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var parentID sql.NullInt64
		err := rows.Scan(&id, &name, &parentID)

		if err != nil {
			log.Fatal(err)
		}

		arr = append(arr, models.ItemCategory{Name: name, ID: id, ParentID: parentID})
	}

	data := models.Response{Err: 0, Message: "", Data: arr}
	writeJSON(w, data)
}

package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"simple-inventory/dao"
	"simple-inventory/models"

	_ "github.com/go-sql-driver/mysql" // Import mysql driver
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

// MainHandler accepts requests pointed to "/"
func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main Handled")
	data := models.Response{Err: 0, Message: []string{}, Data: nil}

	writeJSON(w, data)
}

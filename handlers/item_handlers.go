package handlers

import (
	"encoding/json"
	"net/http"
	"simple-inventory/models"
	"strconv"

	"github.com/gorilla/mux"
)

// GetItemsHandler retrieves the list of all items
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {}

// GetItemHandler handler for retrieving a single item
func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	item, err := itemDao.GetItem(int(id))

	if err != nil {
		data := models.Response{Err: 1, Message: []string{err.Error()}, Data: nil}
		writeJSON(w, data)
		return
	}

	data := models.Response{Err: 0, Message: []string{}, Data: item}
	writeJSON(w, data)
}

// CreateItemHandler creates a new item
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&item)

	err := itemDao.CreateItem(item)

	if err != nil {
		data := models.Response{Err: 1, Message: []string{err.Error()}, Data: nil}
		writeJSON(w, data)
		return
	}
	data := models.Response{Err: 0, Message: []string{}, Data: nil}
	writeJSON(w, data)
}

// SetItemCategoriesHandler assigns categories to an item
func SetItemCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&item)

	err := itemDao.SetCategories(int(id), item.Categories)

	if err != nil {
		var stringErrs []string

		for i := 0; i < len(err); i++ {
			stringErrs = append(stringErrs, err[i].Error())
		}

		data := models.Response{Err: 1, Message: stringErrs, Data: nil}
		writeJSON(w, data)
		return
	}
	data := models.Response{Err: 0, Message: []string{}, Data: nil}
	writeJSON(w, data)
}

// SetItemStockHandler handles requests for updating item stock
func SetItemStockHandler(w http.ResponseWriter, r *http.Request) {
	var stock models.ItemStock
	vars := mux.Vars(r)
	itemID, _ := strconv.ParseInt(vars["id"], 10, 64)
	stockID, _ := strconv.ParseInt(vars["stockID"], 10, 64)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&stock)

	stock.ItemID = int(itemID)
	stock.ID = int(stockID)

	err := itemDao.SetItemStock(stock)

	if err != nil {
		var stringErrs []string

		for i := 0; i < len(err); i++ {
			stringErrs = append(stringErrs, err[i].Error())
		}

		data := models.Response{Err: 1, Message: stringErrs, Data: nil}
		writeJSON(w, data)
		return
	}
	data := models.Response{Err: 0, Message: []string{}, Data: nil}
	writeJSON(w, data)
}

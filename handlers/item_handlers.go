package handlers

import (
	"encoding/json"
	"fmt"
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

	fmt.Println("Get data for itemID ", id)

	item, err := itemDao.GetItem(int(id))
	writeJSON(w, createJSONResponse([]error{err}, item))
}

// CreateItemHandler creates a new item
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&item)

	err := itemDao.CreateItem(item)
	writeJSON(w, createJSONResponse([]error{err}, nil))
}

// SetItemCategoriesHandler assigns categories to an item
func SetItemCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&item)

	err := itemDao.SetCategories(int(id), item.Categories)

	writeJSON(w, createJSONResponse(err, nil))
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

	writeJSON(w, createJSONResponse(err, nil))
}

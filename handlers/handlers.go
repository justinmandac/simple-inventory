package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-inventory/models"
)

// MainHandler accepts requests pointed to "/"
func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main Handled")
	data := models.Response{Err: 0, Message: "ok", Data: nil}
	res, _ := json.Marshal(data)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

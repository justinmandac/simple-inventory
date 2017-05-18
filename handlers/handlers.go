package handlers

import (
	"fmt"
	"net/http"
)

// MainHandler accepts requests pointed to "/"
func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main Handled")
}

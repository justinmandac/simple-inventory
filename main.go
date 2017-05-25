package main

import (
	"log"
	"net/http"
	"simple-inventory/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.MainHandler)
	r.HandleFunc("/api/categories", handlers.GetCategoriesHandler).Methods("GET")
	r.HandleFunc("/api/categories", handlers.CreateCategoryHandler).Methods("POST")
	r.HandleFunc("/api/categories/{id}", handlers.DeleteCategoryHandler).Methods("DELETE")
	r.HandleFunc("/api/categories/{id}", handlers.UpdateCategoryHandler).Methods("PUT")
	r.HandleFunc("/api/categories/{id}", handlers.GetCategoryHandler).Methods("GET")

	r.HandleFunc("/api/items", handlers.GetItemsHandler).Methods("GET")
	r.HandleFunc("/api/items", handlers.CreateItemHandler).Methods("POST")
	r.HandleFunc("/api/items/{id}", handlers.GetItemHandler).Methods("GET")
	r.HandleFunc("/api/items/{id}/categories", handlers.SetItemCategoriesHandler).Methods("PUT")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

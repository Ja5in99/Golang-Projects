package routes

import (
	"github.com/Ja5in99/Golang-Projects/tree/main/Bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.CreateBooks).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}

package main

import (
	"github.com/Ja5in99/Golang-Projects/tree/main/Bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
}

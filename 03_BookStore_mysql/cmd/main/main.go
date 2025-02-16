package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rinerte/01_sharma_go/03_BookStore_mysql/pkg/routes"
	_ "gorm.io/driver/postgres"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

package main

import (
	"log"
	"net/http"

	"github.com/bikashsaud/book_store/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()

	routes.BookStoreRouters(r)

	http.Handle("/", r)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}

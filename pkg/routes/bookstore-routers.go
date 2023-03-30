package routes

import (
	"github.com/bikashsaud/book_store/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookStoreRouters = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

	// book detail
	router.HandleFunc("/book/{bookId}/detail", controllers.GetBookByIdController).Methods("GET")

	// update book
	router.HandleFunc("/book/{bookId}/update", controllers.UpdateBookController).Methods("PUT")
}

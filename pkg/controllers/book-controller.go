package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bikashsaud/book_store/pkg/models"
	"github.com/bikashsaud/book_store/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting book list...")
	newBooks := models.GetBooks()
	response, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	reqVars := mux.Vars(r)
	bookId := reqVars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error parsing book id: ", err)

	}
	bookDetail, _ := models.GetBook(Id)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	// utils.ParseBody(r, CreateBook)
	json.NewDecoder(r.Body).Decode(&CreateBook)
	fmt.Println(CreateBook)

	// data := json.NewDecoder(r.Body).Decode(&CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	fmt.Println(CreateBook)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	reqVars := mux.Vars(r)
	bookId := reqVars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error parsing book: ", err)
	}

	bookDetail, db := models.GetBook(Id)

	if updateBook.Name == "" {
		bookDetail.Name = updateBook.Name
	}
	if updateBook.Author == "" {
		bookDetail.Author = updateBook.Author
	}

	if updateBook.Publication == "" {
		bookDetail.Publication = updateBook.Publication
	}
	db.Save(&bookDetail)

	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

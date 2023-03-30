package models

import (
	"errors"
	"fmt"

	"github.com/bikashsaud/book_store/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {

	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBook(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	if db.RecordNotFound() {
		fmt.Println("Book record not found.")
	} else if db.Error != nil {
		fmt.Println("Error occurred:", db.Error)
	}

	fmt.Println(db, 44444)
	return &getBook, db

}
func GetBookById(id int64) (*Book, error) {
	var book Book
	result := db.Where("ID = ?", id).First(&book)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Book with ID %d not found", id)
		}
		return nil, result.Error
	}
	return &book, nil
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}

func UpdateBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Update(book)
	return book
}

func UpdateBookById(book *Book) error {
	// update book in db
	if err := db.Save(book).Error; err != nil {
		return err
	}
	return nil
}

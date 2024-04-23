package models

import (
	"go-books-api/pkg/config"

	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title        string `json:"title"`
	Author       string `json:"author"`
	Pubulication string `json:"publication"`
}

var db *gorm.DB

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

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("id = ?", id).Find(&book)
	return &book, db
}

func (b *Book) UpdateBook() *Book {
	db.Save(&b)
	return b
}

func DeleteBook(id int) Book {
	var book Book
	db.Where("id = ?", id).Delete(book)
	return book
}

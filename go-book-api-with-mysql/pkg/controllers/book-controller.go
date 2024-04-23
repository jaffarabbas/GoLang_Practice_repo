package controllers

import (
	"encoding/json"
	"fmt"
	"go-books-api/pkg/models"
	"go-books-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var newBook *models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := newBook.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book, _ := models.GetBookById(ID)
	if updateBook.Title != "" {
		book.Title = updateBook.Title
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Pubulication != "" {
		book.Pubulication = updateBook.Pubulication
	}
	book.UpdateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Aman-Shitta/go-bookstore/pkg/models"
	"github.com/Aman-Shitta/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}

	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)

	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()

	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	bookId := params["id"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	params := mux.Vars(r)
	bookId := params["id"]

	bookID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Fprintln(w, "Invalid ID")
	}

	bookDetails, db := models.GetBookById(bookID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	bookID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	book := models.DeleteById(bookID)
	res, _ := json.Marshal(book)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	model "goproject.com/pkg/models"
	util "goproject.com/pkg/utils"
)

// NewBook unexported
var NewBook model.Book

// GetBook exported
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook := model.GetAllBooks()
	res, _ := json.Marshal(newBook); header(w, res)
}

// GetBookID exported
func GetBookID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["bookId"]
	
	ID, err := strconv.ParseInt(bookID, 0,0)
	if err != nil { fmt.Println("error while parsing") }

	bookDetails, _ := model.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails); header(w, res)
}

// CreateBook exported
func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &model.Book{}
	util.ParseBody(r, createBook)

	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook exported
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &model.Book{}
	util.ParseBody(r, updateBook)
	
	params := mux.Vars(r)
	bookID := params["bookId"]
	ID, err := strconv.ParseInt(bookID,0,0)
	if err != nil { fmt.Println("error while parsing") }
	
	bookDetails, db := model.GetBookByID(ID)
	if updateBook.Name != "" { bookDetails.Name = updateBook.Name }
	if updateBook.Author != "" { bookDetails.Author = updateBook.Author }
	if updateBook.Publication != "" { bookDetails.Publication = updateBook.Publication }
	
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	header(w, res)
}

// DeleteBook exported
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["bookId"]

	ID, err := strconv.ParseInt(bookID, 0,0)
	if err != nil { fmt.Println("error while parsing") }
	
	book := model.DeleteBook(ID)
	res, _ := json.Marshal(book); header(w, res)
}

func header(w http.ResponseWriter, res []byte){
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
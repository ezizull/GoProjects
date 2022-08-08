package model

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Book exported
type Book struct{
	gorm.Model
	Name 		string `gorm:"" json:"name"`
	Author 		string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	// config.MySQLConnect()
	// db = config.GetDB()
	// db.AutoMigrate(&Book{})
}

// CreateBook exported
func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllBooks exported
func GetAllBooks() []Book{
	var books []Book
	db.Find(&books)
	return books
}

// GetBookByID exported
func GetBookByID(ID int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

// DeleteBook exported
func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}



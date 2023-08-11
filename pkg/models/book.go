package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Venkatsai50/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name     string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book ,*gorm.DB) {
	var getbook Book
	db:=db.Where("id = ?", id).Find(&getbook)
	return &getbook,db
}


func DeleteBook(id int64) Book{
	var book Book
	db.Where("id = ?", id).Delete(&book)
	return book
}
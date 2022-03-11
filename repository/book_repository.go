package repository

import (
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() []entity.Book
	GetByID(bookID uint64) entity.Book
	GetAllMyBook(userID int64) []entity.Book
	CreateMyBook(b entity.Book) entity.Book
	UpdateMyBook(b entity.Book) entity.Book
	DeleteMyBook(b entity.Book)
}

type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(connection *gorm.DB) BookRepository {
	return &bookConnection{connection: connection}
}

func (db *bookConnection) GetAll() []entity.Book {
	var books []entity.Book
	db.connection.Find(&books)
	return books
}

func (db *bookConnection) GetAllMyBook(userID int64) []entity.Book {
	var books []entity.Book
	db.connection.Where("user_id = ?", userID).Find(&books)
	return books
	// db.connection.Preload("User").Find(&books)
	// return books
}

func (db *bookConnection) GetByID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Where("id = ?", bookID).First(&book)
	return book
}

func (db *bookConnection) CreateMyBook(b entity.Book) entity.Book {
	db.connection.Create(&b)
	return b
}

func (db *bookConnection) UpdateMyBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	return b
}

func (db *bookConnection) DeleteMyBook(b entity.Book) {
	db.connection.Delete(&b)
}

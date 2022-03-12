package repository

import (
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() []entity.Book                  // get all book from database
	GetByID(bookID uint64) entity.Book      // get book by bookID
	GetAllMyBook() []entity.Book            // get all book by userID
	CreateMyBook(b entity.Book) entity.Book // create book by userID
	UpdateMyBook(b entity.Book) entity.Book // update book by userID
	DeleteMyBook(b entity.Book)             // delete book by userID
}

// Create bookConnection struct to implement connection to database
type bookConnection struct {
	connection *gorm.DB // connection to database
}

// NewBookConnection method is used to create a new instance of bookConnection
func NewBookRepository(connection *gorm.DB) BookRepository {
	return &bookConnection{connection: connection}
}

// GetAll method is used to get all book from database
func (db *bookConnection) GetAll() []entity.Book {
	var books []entity.Book                    // create variable books to store all book
	db.connection.Preload("User").Find(&books) // get all book and preload user from book
	return books                               // return all book
}

// GetAllMyBook method is used to get all book by userID
func (db *bookConnection) GetAllMyBook() []entity.Book {
	var books []entity.Book                    // create variable books to store all book
	db.connection.Preload("User").Find(&books) // get all book and preload user from book
	return books                               // return all book
}

// GetByID method is used to get book by bookID
func (db *bookConnection) GetByID(bookID uint64) entity.Book {
	var book entity.Book                              // create variable book
	db.connection.Preload("User").Find(&book, bookID) // get data book from bookID and preload user from book
	return book                                       // return book
}

// CreateMyBook method is used to create book by userID
func (db *bookConnection) CreateMyBook(b entity.Book) entity.Book {
	db.connection.Save(&b)                 // save insert book
	db.connection.Preload("User").Find(&b) // get data user from book
	return b                               // return book
}

// UpdateMyBook method is used to update book by userID
func (db *bookConnection) UpdateMyBook(b entity.Book) entity.Book {
	db.connection.Save(&b)                 // save update book
	db.connection.Preload("User").Find(&b) // get data user from book
	return b                               // return book
}

// DeleteMyBook method is used to delete book by userID
func (db *bookConnection) DeleteMyBook(b entity.Book) {
	db.connection.Delete(&b) // delete book
}

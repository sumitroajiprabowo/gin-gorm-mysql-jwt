package services

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/dto"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/entity"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/repository"
)

type BookService interface {
	GetAll() []entity.Book
	GetByID(bookID uint64) entity.Book
	GetAllMyBook(userID int64) []entity.Book
	CreateMyBook(b dto.BookCreateDTORequest) entity.Book
	UpdateMyBook(b dto.BookUpdateDTORequest) entity.Book
	DeleteMyBook(b entity.Book)
	IsAllowedActionBook(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{bookRepository: bookRepo}
}

func (b *bookService) GetAll() []entity.Book {
	return b.bookRepository.GetAll()
}

func (b *bookService) GetByID(bookID uint64) entity.Book {
	return b.bookRepository.GetByID(bookID)
}

func (b *bookService) GetAllMyBook(userID int64) []entity.Book {
	return b.bookRepository.GetAllMyBook(userID)
}

func (s *bookService) CreateMyBook(b dto.BookCreateDTORequest) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed to map fields %v: ", err)
	}
	result := s.bookRepository.CreateMyBook(book)
	return result
}

func (s *bookService) UpdateMyBook(b dto.BookUpdateDTORequest) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed to map fields %v: ", err)
	}
	result := s.bookRepository.UpdateMyBook(book)
	return result
}

func (s *bookService) DeleteMyBook(b entity.Book) {
	s.bookRepository.DeleteMyBook(b)
}

func (s *bookService) IsAllowedActionBook(userID string, bookID uint64) bool {
	b := s.bookRepository.GetByID(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}

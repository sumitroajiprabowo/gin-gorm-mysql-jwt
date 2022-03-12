package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/dto"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/entity"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/helper"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/services"
)

type BookController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	GetAllMyBook(c *gin.Context)
	CreateMyBook(c *gin.Context)
	UpdateMyBook(c *gin.Context)
	DeleteMyBook(c *gin.Context)
}

type bookController struct {
	bookService services.BookService
	jwtService  services.JWTService
}

func NewBookController(bookServ services.BookService, jwtServ services.JWTService) BookController {
	return &bookController{bookService: bookServ, jwtService: jwtServ}
}

func (c *bookController) GetAll(ctx *gin.Context) {
	var books []entity.Book = c.bookService.GetAll()
	result := helper.SuccessResponse(http.StatusOK, "Get All Data Book", books)
	ctx.JSON(http.StatusOK, result)
}

func (c *bookController) GetByID(ctx *gin.Context) {
	bookID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Book Not Found", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var book entity.Book = c.bookService.GetByID(bookID)
	if book == (entity.Book{}) {
		response := helper.ErrorsResponse(http.StatusNotFound, "Book Not Found Cok", "", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	} else {
		response := helper.SuccessResponse(http.StatusOK, "Get Data Book", book)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *bookController) GetAllMyBook(ctx *gin.Context) {
	var book []entity.Book = c.bookService.GetAllMyBook()
	response := helper.SuccessResponse(http.StatusOK, "Get All Data Book", book)
	ctx.JSON(http.StatusOK, response)
}

func (c *bookController) CreateMyBook(ctx *gin.Context) {

	var bookCreateDTO dto.BookCreateDTORequest
	errDTO := ctx.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Invalid data", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	id, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		bookCreateDTO.UserID = id
	}
	result := c.bookService.CreateMyBook(bookCreateDTO)
	response := helper.SuccessResponse(http.StatusCreated, "Create Data Book", result)
	ctx.JSON(http.StatusCreated, response)

}

func (c *bookController) UpdateMyBook(ctx *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTORequest
	errDTO := ctx.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		result := helper.ErrorsResponse(http.StatusBadRequest, "Invalid data", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedActionBook(userID, bookUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			bookUpdateDTO.UserID = id
		}
		result := c.bookService.UpdateMyBook(bookUpdateDTO)
		response := helper.SuccessResponse(http.StatusOK, "Update Data Book", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.ErrorsResponse(http.StatusForbidden, "Forbidden", "You are not allowed to update this book", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusForbidden, response)
	}
}

func (c *bookController) DeleteMyBook(ctx *gin.Context) {
	var book entity.Book
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Book Not Found", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	book.ID = id
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedActionBook(userID, book.ID) {
		c.bookService.DeleteMyBook(book)
		response := helper.SuccessResponse(http.StatusOK, "Delete Data Book", book)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.ErrorsResponse(http.StatusForbidden, "Forbidden", "You are not allowed to delete this book", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusForbidden, response)
	}
}

func (c *bookController) getUserIDByToken(token string) string {
	myToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := myToken.Claims.(jwt.MapClaims)
	return claims["user_id"].(string)
}

package controllers

import (
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
	// DeleteMyBook(c *gin.Context)
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
	result := c.bookService.GetByID(bookID)
	response := helper.SuccessResponse(http.StatusOK, "Get Data Book By ID", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *bookController) getUserIDByToken(token string) string {
	myToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := myToken.Claims.(jwt.MapClaims)
	return claims["userId"].(string)
}

func (c *bookController) GetAllMyBook(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	id, _ := strconv.ParseUint(userID, 10, 64)
	books := c.bookService.GetAllMyBook(int64(id))
	response := helper.SuccessResponse(http.StatusOK, "Get All Data Book", books)
	ctx.JSON(http.StatusOK, response)
}

func (c *bookController) CreateMyBook(ctx *gin.Context) {

	var bookCreateDTO dto.BookCreateDTORequest
	errDTO := ctx.BindJSON(&bookCreateDTO)
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
	errDTO := ctx.BindJSON(&bookUpdateDTO)
	if errDTO != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Invalid data", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	id, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		bookUpdateDTO.UserID = id
	}
	result := c.bookService.UpdateMyBook(bookUpdateDTO)
	response := helper.SuccessResponse(http.StatusOK, "Update Data Book", result)
	ctx.JSON(http.StatusOK, response)
}
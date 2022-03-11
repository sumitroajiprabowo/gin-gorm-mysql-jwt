package controllers

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/dto"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/helper"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/services"
)

type UserController interface {
	UpdateUser(c *gin.Context)
	GetUser(c *gin.Context)
}

type userController struct {
	userService services.UserService
	jwtService  services.JWTService
}

func NewUserController(userService services.UserService, jwtService services.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTORequest
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.ErrorsResponse(http.StatusUnauthorized, "Failed to process request", errToken.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	userId, err := strconv.ParseUint(claims["userId"].(string), 10, 64)
	if err != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userUpdateDTO.ID = userId
	user := c.userService.UpdateUser(userUpdateDTO)
	response := helper.SuccessResponse(http.StatusOK, "Update User Success", user)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) GetUser(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.ErrorsResponse(http.StatusUnauthorized, "Failed to process request", errToken.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)

	userId, err := strconv.ParseUint(claims["userId"].(string), 10, 64)
	if err != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user := c.userService.GetUser(int64(userId))
	response := helper.SuccessResponse(http.StatusOK, "Get User Success", user)
	ctx.JSON(http.StatusOK, response)
}

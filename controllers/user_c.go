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

// UserController is a struct for user controller
type UserController interface {
	UpdateUser(c *gin.Context) // UpdateUser is a function for update user
	GetUser(c *gin.Context)    // GetUser is a function for get user
}

// userController is a struct for user controller
type userController struct {
	// userService is a new instance of UserService
	userService services.UserService
	// jwtService is a new instance of JWTService
	jwtService services.JWTService
}

// NewUserController is a function for create new instance of UserController
func NewUserController(userService services.UserService, jwtService services.JWTService) UserController {
	return &userController{
		// userService is a new instance of UserService
		userService: userService,
		// jwtService is a new instance of JWTService
		jwtService: jwtService,
	}
}

// UpdateUser is a function for update user
func (c *userController) UpdateUser(ctx *gin.Context) {

	// userUpdateDTO is a new instance of UserUpdateDTORequest
	var userUpdateDTO dto.UserUpdateDTORequest

	// Bind the userUpdateDTO with the request body
	errDTO := ctx.ShouldBind(&userUpdateDTO)

	// Check if there is any error in binding
	if errDTO != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Get the token from the header of the request
	authHeader := ctx.GetHeader("Authorization")

	// Validate the token
	token, errToken := c.jwtService.ValidateToken(authHeader)

	// Check if there is any error in validating token
	if errToken != nil {
		response := helper.ErrorsResponse(http.StatusUnauthorized, "Failed to process request", errToken.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	claims := token.Claims.(jwt.MapClaims) // Get the claims of the token

	// Get the user id from the claims
	userId, err := strconv.ParseUint(claims["user_id"].(string), 10, 64)

	// Check if there is any error in parsing user id
	if err != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userUpdateDTO.ID = userId // Get the user from the database

	user := c.userService.UpdateUser(userUpdateDTO) // Update the user

	response := helper.SuccessResponse(http.StatusOK, "Update User Success", user) // Create the response for the user

	ctx.JSON(http.StatusOK, response) // Return the response
}

// GetUser is a function for get user
func (c *userController) GetUser(ctx *gin.Context) {

	// Get the token from the header of the request
	authHeader := ctx.GetHeader("Authorization")

	// Validate the token
	token, errToken := c.jwtService.ValidateToken(authHeader)

	// Check if there is any error in validating token
	if errToken != nil {
		response := helper.ErrorsResponse(http.StatusUnauthorized, "Failed to process request", errToken.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// Get the claims of the token
	claims := token.Claims.(jwt.MapClaims)

	// Get the user id from the claims
	userId, err := strconv.ParseUint(claims["user_id"].(string), 10, 64)

	// Check if there is any error in parsing user id
	if err != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Get the user from the database with the user id
	user := c.userService.GetUser(int64(userId))

	// Create the response for the user
	response := helper.SuccessResponse(http.StatusOK, "Get User Success", user)

	// Return the response
	ctx.JSON(http.StatusOK, response)
}

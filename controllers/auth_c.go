package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/dto"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/entity"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/helper"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/services"
)

// Auth Controller interface is a contract for all auth controller
type AuthController interface {
	Login(c *gin.Context)    // Login
	Register(c *gin.Context) // Register
}

// Auth Controller struct to implement AuthController interface
type authController struct {
	authService services.AuthService // inject auth service
	jwtService  services.JWTService  // inject jwt service
}

/*
Create a new instance of Auth Controller with auth service and jwt service injected as dependency
*/
func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService, // inject auth service
		jwtService:  jwtService,  // inject jwt service
	}
}

// Login is a function for login
func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTORequest // create new instance of LoginDTORequest

	// bind the loginDTO with the request body
	errDTO := ctx.ShouldBind(&loginDTO)

	// Check if there is any error in binding
	if errDTO != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Check if the email and password is valid
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)

	// Check if the email and password is valid
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.Itoa(int(v.ID)))
		v.Token = generatedToken
		response := helper.SuccessResponse(http.StatusOK, "Login Success", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	// If the email and password is not valid
	response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", "Invalid Credential", helper.EmptyObject{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// Register is a function for register
func (c *authController) Register(ctx *gin.Context) {

	// create new instance of RegisterDTORequest
	var registerDTO dto.RegisterDTORequest

	// bind the registerDTO with the request body
	errDTO := ctx.ShouldBind(&registerDTO)

	// Check if there is any error in binding
	if errDTO != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Check if the email is valid and unique in the database
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.ErrorsResponse(http.StatusConflict, "Failed to process request", "Email already registered", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	} else {
		/*
			if the email is valid and unique in the database then register the user
		*/
		createdUser := c.authService.CreateUser(registerDTO) // create new user

		// generate token
		token := c.jwtService.GenerateToken(strconv.Itoa(int(createdUser.ID)))

		// set token to the user
		createdUser.Token = token

		// response with the user data and token
		response := helper.SuccessResponse(http.StatusCreated, "Register Success", createdUser)

		// return the response
		ctx.JSON(http.StatusCreated, response)
	}
}

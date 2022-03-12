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
	Login(c *gin.Context)
	Register(c *gin.Context)
}

// Auth Controller struct
type authController struct {
	// authService AuthService
	authService services.AuthService
	jwtService  services.JWTService
}

// Create a new instance of Auth Controller
func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTORequest
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.Itoa(int(v.ID)))
		v.Token = generatedToken
		response := helper.SuccessResponse(http.StatusOK, "Login Success", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", "Invalid Credential", helper.EmptyObject{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTORequest
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.ErrorsResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.ErrorsResponse(http.StatusConflict, "Failed to process request", "Email already registered", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.Itoa(int(createdUser.ID)))
		createdUser.Token = token
		response := helper.SuccessResponse(http.StatusCreated, "Register Success", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

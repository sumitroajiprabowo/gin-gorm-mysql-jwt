package controllers

import (
	"github.com/gin-gonic/gin"
)

// Auth Controller interface is a contract for all auth controller
type AuthController interface {
	Login(c *gin.Context)

	Register(c *gin.Context)
}

// Auth Controller struct
type authController struct {
	// authService AuthService
}

// Create a new instance of Auth Controller
func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "login",
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "register",
	})
}

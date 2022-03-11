package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/config"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/controllers"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/repository"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/services"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.SetupDatabase()
	userRepository repository.UserRepository  = repository.NewUserRepository(db)
	jwtService     services.JWTService        = services.NewJWTService()
	userService    services.UserService       = services.NewUserService(userRepository)
	authService    services.AuthService       = services.NewAuthService(userRepository)
	authController                            = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController = controllers.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("/api/user")
	{
		userRoutes.GET("/profile", userController.GetUser)
		userRoutes.PUT("/profile", userController.UpdateUser)
	}

	r.Run()

}

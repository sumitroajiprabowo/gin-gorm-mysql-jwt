package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/config"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/controllers"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/middleware"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/repository"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/services"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.SetupDatabase()
	userRepository repository.UserRepository  = repository.NewUserRepository(db)
	bookRepository repository.BookRepository  = repository.NewBookRepository(db)
	jwtService     services.JWTService        = services.NewJWTService()
	userService    services.UserService       = services.NewUserService(userRepository)
	bookService    services.BookService       = services.NewBookService(bookRepository)
	authService    services.AuthService       = services.NewAuthService(userRepository)
	authController                            = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController = controllers.NewUserController(userService, jwtService)
	bookController controllers.BookController = controllers.NewBookController(bookService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("/api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.GetUser)
		userRoutes.PUT("/profile", userController.UpdateUser)
	}

	bookRoutes := r.Group("api/books", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.GetAll)
		bookRoutes.POST("/", bookController.CreateMyBook)
		bookRoutes.GET("/:id", bookController.GetByID)
		bookRoutes.PUT("/:id", bookController.UpdateMyBook)
	}

	r.Run()

}

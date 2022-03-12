package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/helper"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/services"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization") // Get the token from the header of the request (if any) // Get the token from the header of the request (if any)
		if authHeader == "" {
			response := helper.ErrorsResponse(401, "Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader) // Validate the token
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)             // Get the claims of the token
			log.Println("Claim[user_id]: ", claims["user_id"]) // output the user_id
			log.Println("Claim[issuer] :", claims["issuer"])   // output the issuer
		} else {
			log.Println(err)
			response := helper.ErrorsResponse(401, "Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}

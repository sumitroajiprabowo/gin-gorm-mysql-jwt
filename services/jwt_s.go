package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT Service is a contract of what a JWT Service should be able to do.
type JWTService interface {
	GenerateToken(userId string) string             // Generate a new token
	ValidateToken(token string) (*jwt.Token, error) // Validate the token
}

// jwtCustomClaims is a struct that contains the custom claims for the JWT
type jwtCustomClaim struct {
	UserId             string `json:"userId"` // The userId is the only required field
	jwt.StandardClaims        // This is a standard JWT claim
}

// jwtService is a struct that implements the JWTService interface
type jwtService struct {
	secretKey string // Secret key used to sign the token
	issuer    string // Who creates the token
}

// Create get the secret key from the environment variable
func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "secret"
	}
	return secretKey
}

//NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(), // Call the getSecretKey function to get the secret key
		issuer:    "go-jwt",       // who creates the token
	}
}

// Create a new token object, specifying signing method and the claims
func (s *jwtService) GenerateToken(UserId string) string {

	// Create the Claims struct with the required claims for the JWT
	claims := &jwtCustomClaim{
		UserId, // userId is the only required field
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(), // 1 year expiration
			Issuer:    s.issuer,                           // who creates the token
			IssuedAt:  time.Now().Unix(),                  // when the token was issued/created (now)
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Sign the token with our secret
	t, err := token.SignedString([]byte(s.secretKey))          // Sign the token with an expiration time
	if err != nil {
		panic(err) // If there is an error, panic
	}
	return t // Return the token to the user, along with an expiration time
}

// ValidateToken validates the token and returns the claims
func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	// Parse the token
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok { // Check the signing method
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"]) // Return an error if the signing method isn't HMAC
		}
		return []byte(s.secretKey), nil // Return the key
	})
}

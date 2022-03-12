package services

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/dto"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/entity"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/repository"
)

// Create User Service Interface for User Service Implementation
type UserService interface {
	UpdateUser(user dto.UserUpdateDTORequest) entity.User
	GetUser(userID int64) entity.User
}

// Create userService struct to implement UserService interface
type userService struct {
	userRepository repository.UserRepository
}

// NewUserService method is used to create a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

// UpdateUser method is used to update user
func (s *userService) UpdateUser(user dto.UserUpdateDTORequest) entity.User {
	userToUpdate := entity.User{}                                        // userToUpdate is a new instance of User
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user)) // Fill the userToUpdate with the user data
	if err != nil {
		log.Fatalf("Error while mapping user update dto to entity: %v", err)
	}
	updatedUser := s.userRepository.UpdateUser(userToUpdate) // Update the user
	return updatedUser
}

// GetUser method is used to get user by userID
func (s *userService) GetUser(userID int64) entity.User {
	user := s.userRepository.ProfileUser(userID) // Get the user by userID
	return user
}

package services

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/dto"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/entity"
	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/repository"
)

type UserService interface {
	UpdateUser(user dto.UserUpdateDTORequest) entity.User
	GetUser(userID int64) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (s *userService) UpdateUser(user dto.UserUpdateDTORequest) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Error while mapping user update dto to entity: %v", err)
	}
	updatedUser := s.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (s *userService) GetUser(userID int64) entity.User {
	user := s.userRepository.ProfileUser(userID)
	return user
}

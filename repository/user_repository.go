package repository

import (
	"log"

	"github.com/sumitroajiprabowo/gin-gorm-jwt-mysql/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	//InsertUser is insert user to db
	InsertUser(user entity.User) entity.User

	//UpdateUser is update user to db
	UpdateUser(user entity.User) entity.User

	//VerifyCredential is verify user credential
	VerifyCredential(email string, password string) interface{}

	//IsDuplicateEmail is check duplicate email
	IsDuplicateEmail(email string) (tx *gorm.DB)

	//FindByEmail is find user by email
	FindByEmail(email string) entity.User

	//ProfileUser is find user by id
	ProfileUser(userID string) entity.User
}

//userConnection is a struct that implements connection to db with gorm
type userConnection struct {
	connection *gorm.DB
}

/*
NewUserRepository is creates a new instance of UserRepository with gorm connection instance as parameter and return UserRepository interface instance to use
*/
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

// CreateUser is insert user to db and return user entity to caller function
func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password)) //hash password
	db.connection.Save(&user)                          //save user to db
	return user
}

// UpdateUser is update user to db and return user entity to caller function
func (db *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password)) //hash password
	} else {
		var tempUser entity.User               //get user from db
		db.connection.Find(&tempUser, user.Id) //find user by id
		user.Password = tempUser.Password      //set password to user
	}
	db.connection.Save(&user) //save user to db
	return user
}

// VerifyCredential is verify user credential and return user entity to caller function if credential is correct or return nil if credential is incorrect
func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

// func (db *userConnection) VerifyCredential(email string, password string) interface{} {
// 	var user entity.User
// 	db.connection.Where("email = ?", email).Take(&user)
// 	if user.Id == 0 {
// 		return nil
// 	}
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 	if err != nil {
// 		return nil
// 	}
// 	return user
// }

//IsDuplicateEmail is check duplicate email and return transaction to caller function
func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

// FindByEmail is find user by email and return user entity to caller function
func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

// ProfileUser is find user by id and return user entity to caller function
func (db *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connection.Preload("Books").Preload("Books.User").Find(&user, userID)
	return user
}

// hashAndSalt is hash password and return hashed password
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

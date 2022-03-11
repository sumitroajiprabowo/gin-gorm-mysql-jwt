package dto

// Create Register DTO Request Struct when user register from /register URL
type RegisterDTORequest struct {
	// Name is the name of the user with minimum length of 3 characters and maximum length of 100 characters
	Name string `json:"name" form:"name" binding:"required,min=3,max=100"`
	// Email is the email of the user with regular expression of email format and required
	Email string `json:"email" form:"email" binding:"required,email"`
	// Password is the password of the user with minimum length of 8 characters and maximum length of 100 characters
	Password string `json:"password" form:"password" binding:"required,min=8,max=100"`
}

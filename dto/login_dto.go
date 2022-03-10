package dto

// Create Login DTO Request Struct when user login from /login URL
type LoginDTORequest struct {
	Email    string `json:"email" form:"email" binding:"required, email"`
	Password string `json:"password" form:"password" binding:"required,min=8"`
}

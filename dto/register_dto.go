package dto

// Create Register DTO Request Struct when user register from /register URL
type RegisterDTORequest struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min=1"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min=8"`
}

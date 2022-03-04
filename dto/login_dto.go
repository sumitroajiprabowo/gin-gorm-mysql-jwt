package dto

// Create Login DTO Request Struct when user login from /login URL
type LoginDTORequest struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omnitempty" form:"password,omnitempty" binding:"required"`
}

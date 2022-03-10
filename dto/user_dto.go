package dto

// Create User Update DTO Request Struct when user update profile
type UserUpdateDTORequest struct {
	Id       uint64 `json:"id" form:"id"`
	Email    string `json:"email" form:"email" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

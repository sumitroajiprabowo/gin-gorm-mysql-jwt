package dto

// Create Book Update DTO Request when user update book
type BookUpdateDTORequest struct {
	ID          uint64 `json:"id" form:"id"`
	Title       string `json:"title" form:"title" binding:"required"`
	Author      string `json:"author" form:"author" binding:"required"`
	Price       int64  `json:"price" form:"price" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omnitempty" form:"user_id,omitempty"`
}

// Create Book Create DTO Request when user create book
type BookCreateDTORequest struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Author      string `json:"author" form:"author" binding:"required"`
	Price       int64  `json:"price" form:"price" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omnitempty" form:"user_id,omitempty"`
}

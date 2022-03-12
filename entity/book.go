package entity

// Create Book struct representing the book table in the database
type Book struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"` // Primary key, auto-increment id with json tag id for json marshalling
	Title       string `gorm:"type:varchar(255)" json:"title"`       // Data type varchar with json tag name for json marshalling
	Author      string `gorm:"type:varchar(255)" json:"author"`      // Data type varchar with json tag name for json marshalling
	Price       int64  `gorm:"type:int(11)" json:"price"`            // Data type varchar with json tag name for json marshalling
	Description string `gorm:"type:varchar(255)" json:"description"` // Data type varchar with json tag name for json marshalling
	UserID      uint64 `gorm:"not_null" json:"-"`
	// Create a foreign key to user table with json tag user for json marshalling
	User *User `gorm:"foreignkey:UserID;references:ID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}

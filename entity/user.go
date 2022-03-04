package entity

// Create User struct representing the user table in the database
type User struct {
	Id       int64  `gorm:"primary_key;auto_increment" json:"id"`        // Primary key, auto-increment id with json tag id for json marshalling
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"` // Unique index for email with json tag email for json marshalling
	Name     string `gorm:"type:varchar(255)" json:"name"`               // Data type varchar with json tag name for json marshalling
	Password string `gorm:"->;<-;not null" json:"-"`                     // Password field with json tag password for json marshalling
	Token    string `gorm:"-" json:"token,omitempty"`                    // Token field with json tag token for json marshalling
}

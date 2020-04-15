package models

// User Model
//
// User хранит информацию о пользователях системы
//
type User struct {
	// the user ID
	ID uint `gorm:"primary_key;unique_index;AUTO_INCREMENT" json:"id"`
	// the user token
	Token string `gorm:"type:varchar(180);unique_index" json:"token"`
	Login string
	Pass  string
	Role  string // 0-guest, 1-user, 2- ..., 256 - "admin"
	// the user name or nick
	Name string `gorm:"type:text" form:"name" query:"name" json:"name"`
}

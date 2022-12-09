package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Threads  []Thread  `json:"threads"`
	Comments []Comment `json:"comments"`
	Username string    `json:"username" gorm:"unique"`
	Password string    `json:"password"`
	//email address, profile photo (relative url) ?, karma?
}

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Threads  []Thread  `json:"threads"`
	Comments []Comment `json:"comments"`
	Username string    `json:"username" gorm:"unique"`
	Password []byte    `json:"-"`
	//email address, profile photo (relative url) ?, karma?
}

type Password struct {
	Password string `json:"password"`
}

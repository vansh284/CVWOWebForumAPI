package models

import "gorm.io/gorm"

/*
Two separate representations for password, one in the User struct, one in the Password
struct. The User struct is meant to be sent to the client, hence the password
has no json representation. The Password struct is meant to be sent to the database, hence
the password has a json representation.
*/

type User struct {
	gorm.Model
	Threads  []Thread  `json:"threads"`
	Comments []Comment `json:"comments"`
	Username string    `json:"username" gorm:"unique"`
	Password []byte    `json:"-"`
}

type Password struct {
	Password string `json:"password"`
}

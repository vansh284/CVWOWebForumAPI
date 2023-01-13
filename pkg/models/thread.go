package models

import (
	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Tag      string    `json:"tag"`
	Author   string    `json:"author"`
	Image    string    `json:"image"`
	Comments []Comment `json:"comments"`
	UserID   uint      `json:"user_id"`
}

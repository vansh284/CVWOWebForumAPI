package models

import (
	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	Content  string    `json:"content"`
	Tag      string    `json:"tag"`
	Comments []Comment `json:"comments"`
	UserID   uint      `json:"user_id"`
}

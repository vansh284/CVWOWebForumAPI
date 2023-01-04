package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content  string `json:"content"`
	Author   string `json:"author"`
	UserID   uint   `json:"user_id"`
	ThreadID uint   `json:"thread_id"`
}

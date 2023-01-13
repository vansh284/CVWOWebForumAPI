package utils

import (
	"errors"

	"github.com/vansh284/CVWOWebForumAPI/pkg/config"
	"github.com/vansh284/CVWOWebForumAPI/pkg/models"
	"gorm.io/gorm"
)

func FindUserByName(username string, user *models.User) error {
	// Helper function that finds the user with given username from the database
	db := config.GetDB()
	if res := db.Where("username = ?", username).First(&user); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	return nil
}

func FindUserByID(id int, user *models.User) error {
	// Helper function that finds the user with given id from the database
	db := config.GetDB()
	if res := db.Find(&user, id); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	return nil
}

func FindThreadByID(id int, thread *models.Thread) error {
	// Helper function that finds the thread with given id from the database
	db := config.GetDB()
	if res := db.Find(&thread, id); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	return nil
}

func FindCommentByID(id int, comment *models.Comment) error {
	// Helper function that finds the comment with given id from the database
	db := config.GetDB()
	if res := db.Find(&comment, id); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	return nil
}

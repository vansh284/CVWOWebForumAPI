package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/vansh284/CVWOWebForumAPI/pkg/config"
	"github.com/vansh284/CVWOWebForumAPI/pkg/models"
	"gorm.io/gorm"
)

func findCommentByID(id int, comment *models.Comment) error {
	// Helper function that finds the comment with given id from the database
	db := config.GetDB()
	if res := db.Find(&comment, id); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	return nil
}

func GetCommentsT(c *fiber.Ctx) error {
	db := config.GetDB()
	var comments []models.Comment
	threadID, err := c.ParamsInt("thread_id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := db.Where("thread_id = ?", threadID).Find(&comments).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(comments)
}

func CreateComment(c *fiber.Ctx) error {
	db := config.GetDB()
	var comment models.Comment
	var user models.User
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findUserByID(int(comment.UserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	user.Comments = append(user.Comments, comment)
	if err := db.Save(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(user.Comments[len(user.Comments)-1])
}

func EditComment(c *fiber.Ctx) error {
	db := config.GetDB()
	var oldComment models.Comment
	var newComment models.Comment
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := c.BodyParser(&newComment); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findCommentByID(id, &oldComment); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	oldComment.Content = newComment.Content
	if err := db.Save(&oldComment).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(oldComment)
}

func DeleteComment(c *fiber.Ctx) error {
	db := config.GetDB()
	var comment models.Comment
	var thread models.Thread
	var user models.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findCommentByID(id, &comment); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findUserByID(int(comment.UserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findThreadByID(int(comment.ThreadID), &thread); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := db.Model(&user).Association("Comments").Delete(&comment); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := db.Model(&thread).Association("Comments").Delete(&comment); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := db.Delete(&comment).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(comment)
}

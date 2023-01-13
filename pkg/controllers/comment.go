package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vansh284/CVWOWebForumAPI/pkg/config"
	"github.com/vansh284/CVWOWebForumAPI/pkg/models"
	"github.com/vansh284/CVWOWebForumAPI/pkg/utils"
)

func GetCommentsT(c *fiber.Ctx) error {
	db := config.GetDB()
	var comments []models.Comment
	if _, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if threadID, err := c.ParamsInt("thread_id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := db.Where("thread_id = ?", threadID).Find(&comments).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.JSON(comments)
}

func CreateComment(c *fiber.Ctx) error {
	db := config.GetDB()
	var comment models.Comment
	var user models.User
	if userID, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := utils.FindUserByID(userID, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else {
		user.Comments = append(user.Comments, comment)
		if err := db.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
	}
	return c.JSON(user.Comments[len(user.Comments)-1])
}

func EditComment(c *fiber.Ctx) error {
	db := config.GetDB()
	var oldComment models.Comment
	var newComment models.Comment
	if userID, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if id, err := c.ParamsInt("id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := c.BodyParser(&newComment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := utils.FindCommentByID(id, &oldComment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if oldComment.UserID != uint(userID) {
		return c.Status(fiber.StatusUnauthorized).JSON("Not allowed to edit other users' comments")
	} else {
		oldComment.Content = newComment.Content
		if err := db.Save(&oldComment).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
	}
	return c.JSON(oldComment)
}

func DeleteComment(c *fiber.Ctx) error {
	db := config.GetDB()
	var (
		comment models.Comment
		thread  models.Thread
		user    models.User
	)
	if userID, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if id, err := c.ParamsInt("id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := utils.FindCommentByID(id, &comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if comment.UserID != uint(userID) {
		return c.Status(fiber.StatusUnauthorized).JSON("Not allowed to delete other users' comments")
	} else if err := utils.FindUserByID(userID, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := utils.FindThreadByID(int(comment.ThreadID), &thread); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := db.Delete(&comment).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.JSON(comment)
}

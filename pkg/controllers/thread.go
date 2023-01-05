package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vansh284/CVWOWebForumAPI/pkg/config"
	"github.com/vansh284/CVWOWebForumAPI/pkg/models"
	"github.com/vansh284/CVWOWebForumAPI/pkg/utils"
)

func GetThreads(c *fiber.Ctx) error {
	fmt.Println("getthreads")
	db := config.GetDB()
	var threads []models.Thread
	if _, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if err := db.Find(&threads).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.JSON(threads)
}

func CreateThread(c *fiber.Ctx) error {
	db := config.GetDB()
	var thread models.Thread
	var user models.User
	if userID, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if err := c.BodyParser(&thread); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := utils.FindUserByID(userID, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else {
		user.Threads = append(user.Threads, thread)
		if err := db.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
	}
	return c.JSON(user.Threads[len(user.Threads)-1])
}

func EditThread(c *fiber.Ctx) error {
	db := config.GetDB()
	var oldThread models.Thread
	var newThread models.Thread
	if userID, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if id, err := c.ParamsInt("id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := c.BodyParser(&newThread); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := utils.FindThreadByID(id, &oldThread); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if oldThread.UserID != uint(userID) {
		return c.Status(fiber.StatusUnauthorized).JSON("Not allowed to edit other users' threads")
	} else {
		oldThread.Content = newThread.Content
		oldThread.Tag = newThread.Tag
		oldThread.Title = newThread.Title
		oldThread.Image = newThread.Image
		if err := db.Save(&oldThread).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
	}
	return c.JSON(oldThread)
}

func DeleteThread(c *fiber.Ctx) error {
	db := config.GetDB()
	var thread models.Thread
	var user models.User
	if userID, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if id, err := c.ParamsInt("id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := utils.FindThreadByID(id, &thread); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if thread.UserID != uint(userID) {
		return c.Status(fiber.StatusUnauthorized).JSON("Not allowed to delete other users' threads")
	} else if err := utils.FindUserByID(userID, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := db.Delete(&thread).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := db.Where("thread_id = ?", id).Delete(&models.Comment{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.JSON(thread)
}

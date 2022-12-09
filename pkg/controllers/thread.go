package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/vansh284/CVWOWebForumAPI/pkg/config"
	"github.com/vansh284/CVWOWebForumAPI/pkg/models"
	"gorm.io/gorm"
)

func findThreadByID(id int, thread *models.Thread) error {
	// Helper function that finds the thread with given id from the database
	db := config.GetDB()
	if res := db.Find(&thread, id); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	return nil
}

func GetThreads(c *fiber.Ctx) error {
	db := config.GetDB()
	var threads []models.Thread
	if err := db.Find(&threads).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(threads)
}

func CreateThread(c *fiber.Ctx) error {
	db := config.GetDB()
	var thread models.Thread
	var user models.User
	if err := c.BodyParser(&thread); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findUserByID(int(thread.UserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	user.Threads = append(user.Threads, thread)
	if err := db.Save(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(user.Threads[len(user.Threads)-1])
}

func EditThread(c *fiber.Ctx) error {
	db := config.GetDB()
	var oldThread models.Thread
	var newThread models.Thread
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := c.BodyParser(&newThread); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findThreadByID(id, &oldThread); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	oldThread.Content = newThread.Content
	oldThread.Tag = newThread.Tag
	if err := db.Save(&oldThread).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(oldThread)
}

func DeleteThread(c *fiber.Ctx) error {
	db := config.GetDB()
	var thread models.Thread
	var user models.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findThreadByID(id, &thread); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findUserByID(int(thread.UserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := db.Model(&user).Association("Threads").Delete(&thread); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := db.Delete(&thread).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(thread)
}

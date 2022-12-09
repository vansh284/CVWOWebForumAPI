package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/vansh284/CVWOWebForumAPI/pkg/config"
	"github.com/vansh284/CVWOWebForumAPI/pkg/models"
	"gorm.io/gorm"
)

func findUserByName(username string, user *models.User) error {
	// Helper function that finds the user with given username from the database
	db := config.GetDB()
	if res := db.Where("username = ?", username).First(&user); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	return nil
}

func findUserByID(id int, user *models.User) error {
	// Helper function that finds the user with given id from the database
	db := config.GetDB()
	if res := db.Find(&user, id); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	username := c.Params("username")
	if err := findUserByName(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	db := config.GetDB()
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	db.Create(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	db := config.GetDB()
	var user models.User
	username := c.Params("username")
	if err := findUserByName(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := db.Delete(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(user)
}

func UpdateUserUsername(c *fiber.Ctx) error {
	type UserUsername struct {
		Username string `json:"username"`
	}
	db := config.GetDB()
	var user models.User
	username := c.Params("username")
	var userUsername UserUsername
	if err := findUserByName(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := c.BodyParser(&userUsername); err != nil {
		//Known that doesn't cause an error when the JSON provided cannot be binded to the struct, it just returns an empty string.
		//Only causes error if the JSON is invalid.
		return c.Status(400).JSON(err.Error())
	}
	//Add check to make sure username is not empty (will probably be part of some validation helper somwhere)
	if err := db.Model(&user).Update("username", userUsername.Username).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(user)
}

func UpdateUserPassword(c *fiber.Ctx) error {
	type UserPassword struct {
		Password string `json:"password"`
	}
	db := config.GetDB()
	var user models.User
	username := c.Params("username")
	var userPassword UserPassword
	if err := findUserByName(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := c.BodyParser(&userPassword); err != nil {
		//Known that doesn't cause an error when the JSON provided cannot be binded to the struct, it just returns an empty string.
		//Only causes error if the JSON is invalid.
		return c.Status(400).JSON(err.Error())
	}
	//Add check to make sure username is not empty (will probably be part of some validation helper somwhere)
	if err := db.Model(&user).Update("password", userPassword.Password).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(user)
}

func UpdateUserThreads(c *fiber.Ctx) error {
	type NewThread struct {
		Thread models.Thread `json:"thread"` //note that the json field here is 'thread' and not 'threads'
	}
	db := config.GetDB()
	var user models.User
	username := c.Params("username")
	var newThread NewThread
	if err := findUserByName(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := c.BodyParser(&newThread); err != nil {
		//Known that doesn't cause an error when the JSON provided cannot be binded to the struct, it just returns an empty string.
		//Only causes error if the JSON is invalid.
		return c.Status(400).JSON(err.Error())
	}
	//Add check to make sure username is not empty (will probably be part of some validation helper somwhere)
	if err := db.Model(&user).Update("threads", append(user.Threads, newThread.Thread)).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.JSON(user)
}

func UpdateUserComments(c *fiber.Ctx) error {
	return c.JSON("Comments Updated!")
}

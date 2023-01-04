package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vansh284/CVWOWebForumAPI/pkg/config"
	"github.com/vansh284/CVWOWebForumAPI/pkg/models"
	"github.com/vansh284/CVWOWebForumAPI/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c *fiber.Ctx) error {
	var user models.User
	if id, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if err := utils.FindUserByID(id, &user); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	db := config.GetDB()
	var (
		user     models.User
		password models.Password
	)
	if err := c.BodyParser(&password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if password, err := bcrypt.GenerateFromPassword([]byte(password.Password), 10); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else {
		user.Password = password
		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
	}
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var (
		checkUser, realUser models.User
		checkPassword       models.Password
	)
	if err := c.BodyParser(&checkPassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := c.BodyParser(&checkUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := utils.FindUserByName(checkUser.Username, &realUser); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	} else if err := bcrypt.CompareHashAndPassword(realUser.Password, []byte(checkPassword.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := utils.GenerateJWT(c, int(realUser.ID)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.JSON(realUser)
}

func Logout(c *fiber.Ctx) error {
	if _, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}
	utils.ExpireCookie(c)
	return c.JSON("Logged Out")
}

func DeleteUser(c *fiber.Ctx) error {
	db := config.GetDB()
	var user models.User
	if id, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if err := utils.FindUserByID(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := db.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := db.Where("user_id = ?", id).Delete(&models.Thread{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := db.Where("user_id = ?", id).Delete(&models.Comment{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	utils.ExpireCookie(c)
	return c.JSON(user)
}

func UpdateUserUsername(c *fiber.Ctx) error {
	db := config.GetDB()
	var user, oldUser models.User
	if id, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if err := c.BodyParser(&oldUser); err != nil { //Known that doesn't cause an error when the JSON provided cannot be binded to the struct, it just returns an empty string. Only causes error if the JSON is invalid.
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := utils.FindUserByID(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := db.Model(&user).Update("username", oldUser.Username).Error; err != nil { //Add check to make sure username is not empty (will probably be part of some validation helper somwhere)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.JSON(user)
}

func UpdateUserPassword(c *fiber.Ctx) error {
	db := config.GetDB()
	var (
		user     models.User
		password models.Password
	)
	if id, err := utils.ValidateJWT(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	} else if err := c.BodyParser(&password); err != nil { //Known that doesn't cause an error when the JSON provided cannot be binded to the struct, it just returns an empty string. Only causes error if the JSON is invalid.
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := utils.FindUserByID(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if password, err := bcrypt.GenerateFromPassword([]byte(password.Password), 10); err != nil { //Add check to make sure password is not empty (will probably be part of some validation helper somwhere)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else if err := db.Model(&user).Update("password", password).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.JSON(user)
}

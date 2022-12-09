package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vansh284/CVWOWebForumAPI/pkg/config"
	"github.com/vansh284/CVWOWebForumAPI/pkg/models"
	"github.com/vansh284/CVWOWebForumAPI/pkg/routes"
)

func main() {
	app := fiber.New()
	routes.InitThreadRoutes(app)
	config.ConnectDB()
	db := config.GetDB()
	db.AutoMigrate(&models.User{}, &models.Thread{}, &models.Comment{})
	app.Listen(":3000")
}

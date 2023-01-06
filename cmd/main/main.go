package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vansh284/CVWOWebForumAPI/pkg/config"
	"github.com/vansh284/CVWOWebForumAPI/pkg/models"
	"github.com/vansh284/CVWOWebForumAPI/pkg/routes"
)

func main() {
	app := fiber.New()
	config.SetUpEnvironment()
	envMap := config.GetEnvMap()
	config.ConnectDB()
	db := config.GetDB()
	db.AutoMigrate(&models.User{}, &models.Thread{}, &models.Comment{})
	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	routes.InitThreadRoutes(app)
	app.Listen(envMap["LOCAL_PORT"])
}

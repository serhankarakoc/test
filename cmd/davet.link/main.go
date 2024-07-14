package main

import (
	"log"

	"davet.link/configs"
	"davet.link/middlewares"
	"davet.link/routers"
	"davet.link/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/django/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.InitLogger()
	defer utils.GetLogger().Sync()

	configs.SetupSession()

	engine := django.New("views", ".django")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")

	store := session.New()

	app.Use(middlewares.SessionMiddleware())
	app.Use(middlewares.CsrfMiddleware())
	/* app.Use(middlewares.ContextMiddleware) */

	routers.SetupRouters(app, store)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"fiber-api/database"
	"fiber-api/handlers"
	"fiber-api/models"
	"log"

	_ "fiber-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})

	app := fiber.New()

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":3000"))
}

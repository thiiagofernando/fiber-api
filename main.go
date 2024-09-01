package main

import (
	database "fiber-api/db"
	"fiber-api/handlers"
	"fiber-api/models"
	"log"

	_ "fiber-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	// Conectar ao banco de dados
	database.ConnectDB()

	// Aplicar migrações automaticamente
	database.DB.AutoMigrate(&models.User{})

	// iniciar o servidor Fiber
	app := fiber.New()

	// Definir rotas e handlers aqui
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":3000"))
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"go-web-service-demo/database"
	"go-web-service-demo/internal/handler/user"
	"go-web-service-demo/internal/middleware"
	"log"
)

func main() {
	app := fiber.New()
	api := app.Group("v1", middleware.Logging)
	users := api.Group("users")
	users.Get("/", userHandler.GetAll)
	users.Get("/:id", userHandler.Get)
	users.Post("/", userHandler.Create)
	users.Put("/:id", userHandler.Update)
	users.Delete("/:id", userHandler.Delete)

	database.ConnectDB()

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
		return
	}
}

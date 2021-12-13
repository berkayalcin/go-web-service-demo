package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"go-web-service-demo/database"
	_ "go-web-service-demo/docs"
	"go-web-service-demo/internal/handler/user"
	"go-web-service-demo/internal/middleware"
	"log"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	app := fiber.New()
	api := app.Group("v1", middleware.Logging)
	users := api.Group("users")
	app.Get("/swagger/*", swagger.Handler)
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

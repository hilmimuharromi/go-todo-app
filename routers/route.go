package routers

import (
	"github.com/gofiber/fiber/v2"
	"todo-app/handlers"
)

func SetupRouters() {
	app := fiber.New()
	app.Get("/ping", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{"data": "pong test post staging"})
		return nil
	})

	app.Get("/health-check", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{"data": "server is running"})
		return nil
	})
	api := app.Group("/api/v1")

	api.Post("/login", handlers.Login)
	api.Post("/register", handlers.RegisterUser)

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}

}

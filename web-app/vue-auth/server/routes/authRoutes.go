package routes

import (
	"go-auth/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/api/register", handler.Register)
	app.Post("/api/login", handler.Login)
}

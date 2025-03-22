package main

import (
	"go-auth/database"
	"go-auth/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	// database connection

	if _, err := database.ConnectDB(); err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}
	routes.SetupAuthRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

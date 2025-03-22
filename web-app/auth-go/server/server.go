package main

import (
	"gofiber-auth/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"golang.org/x/crypto/bcrypt"
)

// Login Struct
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// register Struct
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// function to hash password
func hashPass(password string) (string, error) {
	//
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
func main() {
	app := fiber.New()
	// use cors
	app.Use(cors.New())

	// connect to database
	database.ConnectDB()
	// login endpoint
	app.Post("/api/login", func(c *fiber.Ctx) error {
		// initialze login struct inside of login endpoint
		var loginData LoginRequest

		// validate the incoming request based on initialized struct
		if err := c.BodyParser(&loginData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request",
			})
		}

		// display the data
		return c.JSON(fiber.Map{
			"message":  "Login data received",
			"email":    loginData.Email,
			"password": loginData.Password,
		})

	})

	// register endpoint
	app.Post("/api/register", func(c *fiber.Ctx) error {
		var registerData RegisterRequest

		// register data validation
		if err := c.BodyParser(&registerData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request",
			})
		}
		// hash the password
		hashedPassword, err := hashPass(registerData.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
		}
		// insert new data from database
		_, err = database.DB.Exec(
			"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
			registerData.Username, registerData.Email, hashedPassword,
		)
		return c.JSON(fiber.Map{"message": "User registered successfully"})

	})
	app.Listen(":3000")
}

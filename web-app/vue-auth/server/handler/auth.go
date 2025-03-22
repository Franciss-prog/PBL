package handler

import (
	"context"
	"fmt"
	"go-auth/database"
	"go-auth/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// function to hash Password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}
func authHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}

// function to register a new user
func Register(c *fiber.Ctx) error {
	var RegisterData models.User

	// parsed the data coming for client
	if err := c.BodyParser(&RegisterData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data didnt recieved",
		})
	}
	// form validtation
	if RegisterData.Email == "" || RegisterData.Username == "" || RegisterData.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data didnt recieved",
		})
	}
	// hash the password
	hashedPassword, err := HashPassword(RegisterData.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash the password",
		})
	}
	// initialized a new value based on Model
	RegisterData.Password = hashedPassword
	RegisterData.CreatedAt = time.Now()

	// insert to user creds to database
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := collection.InsertOne(ctx, RegisterData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}
	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}

// function to handle login
func Login(c *fiber.Ctx) error {
	var LoginData models.User

	if err := c.BodyParser(&LoginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data didnt recieved",
		})

	}

	// form validtation
	if LoginData.Email == "" || LoginData.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data didnt recieved",
		})
	}

	// find the user to  database
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User

	if err := collection.FindOne(ctx, bson.M{"email": LoginData.Email}).Decode(&user); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	if !authHash(LoginData.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}
	message := fmt.Sprintf("Welcome to my App %s", user.Username)
	return c.Status(201).JSON(fiber.Map{"message": message})
}

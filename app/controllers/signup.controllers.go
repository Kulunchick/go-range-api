package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kulunchick/go-range-api/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func registerUser(user *models.User, users *mongo.Collection) {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	users.InsertOne(context.TODO(), bson.M{
		"username": user.Username,
		"email":    user.Email,
		"password": string(password),
	})
}

func signUp(c *fiber.Ctx) error {
	input := new(models.User)
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}
	db := c.Locals("DB").(*mongo.Database)
	users := db.Collection("users")

	value, err := users.CountDocuments(context.TODO(), bson.M{"email": input.Email})
	//mongo: no documents in result
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	if value >= 1 {
		return c.Status(409).JSON(fiber.Map{
			"message": "User already exists",
		})
	} else {
		go registerUser(input, users)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created",
	})
}

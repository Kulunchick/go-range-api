package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/kulunchick/go-range-api/app/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupRoutes(app *fiber.App) {
	routes.AuthRoute(app.Group("/"))
}

func connectToDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(os.Getenv("MONGO_DB"))
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	DB := connectToDB()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("DB", DB)
		return c.Next()
	})

	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}

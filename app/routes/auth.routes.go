package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(route fiber.Router) {
	route.Post("/signup", controllers.signUp)
}

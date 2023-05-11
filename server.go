package main

import (
	"github.com/MrRytis/go-fiber-test/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.NewRouter(app)

	app.Listen(":3000")
}

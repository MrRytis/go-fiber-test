package router

import (
	"github.com/MrRytis/go-fiber-test/handler"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App) {
	app.Get("/api/ping", handler.Ping)
	app.Post("/api/auth/register", handler.Register)
}

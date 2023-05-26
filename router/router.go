package router

import (
	"github.com/MrRytis/go-fiber-test/handler"
	"github.com/MrRytis/go-fiber-test/middleware"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, application *handler.App, middleware *middleware.Middleware) {
	// API v1
	v1 := app.Group("/api/v1")

	// Auth
	auth := v1.Group("/auth", middleware.RateLimiter())
	auth.Post("/register", application.Register)
	auth.Post("/login", application.Login)
	auth.Post("/verify", application.Verify)
	auth.Post("/reminder", application.Reminder)
	auth.Post("/reset", application.Reset)
	auth.Post("/refresh", application.Refresh)
	auth.Post("/logout", middleware.Auth, application.Logout)

	// User
	user := v1.Group("/user", middleware.Auth)
	user.Get("/details", application.GetUserDetails)
	user.Put("/details", application.UpdateUserDetails)
	user.Delete("/delete", application.DeleteUser)
}

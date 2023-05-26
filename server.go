package main

import (
	"github.com/MrRytis/go-fiber-test/exception"
	"github.com/MrRytis/go-fiber-test/handler"
	"github.com/MrRytis/go-fiber-test/middleware"
	"github.com/MrRytis/go-fiber-test/router"
	"github.com/MrRytis/go-fiber-test/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	app.Use(recover.New())

	db := utils.NewDb()
	cache := utils.NewCache()

	handlerApp := handler.NewApp(db, cache)
	middlewareApp := middleware.NewMiddleware(db, cache)

	router.NewRouter(app, handlerApp, middlewareApp)

	log.Fatal(app.Listen(":3000"))
}

package exception

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Error()
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return ctx.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": message,
	})
}

package handler

import (
	"github.com/MrRytis/go-fiber-test/model"
	"github.com/MrRytis/go-fiber-test/request"
	"github.com/MrRytis/go-fiber-test/response"
	"github.com/MrRytis/go-fiber-test/service"
	"github.com/gofiber/fiber/v2"
)

func (app App) GetUserDetails(c *fiber.Ctx) error {
	uid := c.Locals("user").(string)

	var user model.User
	err := app.db.Where("uid = ?", uid).First(user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(response.User{
		Uid:     user.Uid,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Icon:    user.Icon,
	})
}

func (app App) UpdateUserDetails(c *fiber.Ctx) error {
	req := new(request.UpdateUserDetailsRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	uid := c.Locals("user").(string)
	var user model.User
	err := app.db.Where("uid = ?", uid).First(&user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	user.Name = req.Name
	user.Surname = req.Surname

	err = app.db.Save(&user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(response.User{
		Uid:     user.Uid,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Icon:    user.Icon,
	})
}

func (app App) DeleteUser(c *fiber.Ctx) error {
	uid := c.Locals("user").(string)

	var user model.User
	err := app.db.Where("uid = ?", uid).Delete(user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	service.BlackListToken(app.cache, c.Locals("jwt").(string), c.Locals("exp").(int64))

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

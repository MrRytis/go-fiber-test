package handler

import (
	"github.com/MrRytis/go-fiber-test/model"
	"github.com/MrRytis/go-fiber-test/request"
	"github.com/MrRytis/go-fiber-test/response"
	"github.com/MrRytis/go-fiber-test/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

func (app App) Register(c *fiber.Ctx) error {
	req := new(request.RegisterRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	hashedPassword, err := service.HashPassword(req.Password)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	user := model.User{
		Uid:           uuid.New().String(),
		Email:         req.Email,
		Password:      hashedPassword,
		Name:          req.Name,
		Surname:       req.Surname,
		Icon:          nil,
		EmailVerified: false,
	}

	err = app.db.Create(&user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	service.SendMail(user.Email) //TODO: Implement mail sending

	return c.Status(fiber.StatusCreated).JSON(response.AuthUser{
		UserId:  user.Uid,
		Message: "User created",
	})
}

func (app App) Login(c *fiber.Ctx) error {
	req := new(request.LoginRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	user := model.User{}
	err := app.db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if service.CheckUserPassword(req.Password, user.Password) != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Email or password is incorrect")
	}

	if !user.EmailVerified {
		return fiber.NewError(fiber.StatusForbidden, "Email is not verified")
	}

	jwt, err := service.GenerateJWT(user)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	refreshJWT, err := service.GenerateRefreshToken(user)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(response.Login{
		AccessToken:  jwt,
		RefreshToken: refreshJWT,
		ExpiresAt:    time.Now().Add(service.AccessTokenJwtExpDuration).Format(time.RFC3339),
	})
}

func (app App) Verify(c *fiber.Ctx) error {
	req := new(request.VerifyRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	user := model.User{}
	err := app.db.Where("uid = ?", req.UserId).First(&user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	user.EmailVerified = true
	err = app.db.Save(&user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	jwt, err := service.GenerateJWT(user)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	refreshJWT, err := service.GenerateRefreshToken(user)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(response.Login{
		AccessToken:  jwt,
		RefreshToken: refreshJWT,
		ExpiresAt:    time.Now().Add(service.AccessTokenJwtExpDuration).Format(time.RFC3339),
	})
}

func (app App) Reminder(c *fiber.Ctx) error {
	req := new(request.ReminderRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	user := model.User{}
	err := app.db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	user.Reset = true
	err = app.db.Save(&user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	service.SendMail(user.Email) //TODO: Implement mail sending

	return c.Status(fiber.StatusOK).JSON(response.AuthUser{
		UserId:  user.Uid,
		Message: "Reminder sent",
	})
}

func (app App) Reset(c *fiber.Ctx) error {
	req := new(request.ResetRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	user := model.User{}
	err := app.db.Where("uid = ? & reset = true", req.UserId).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "User not found or reset token is invalid")
		}

		return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	pass, err := service.HashPassword(req.Password)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	user.Password = pass
	err = app.db.Save(&user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(response.AuthUser{
		UserId:  user.Uid,
		Message: "Password changed",
	})
}

func (app App) Refresh(c *fiber.Ctx) error {
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	if tokenString == "" {
		return fiber.ErrUnauthorized
	}

	claims, err := service.ParseJWT(tokenString)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	if service.IsBlacklisted(app.cache, tokenString) {
		return fiber.ErrUnauthorized
	}

	uid := claims["uid"].(string)
	user := model.User{}
	err = app.db.Where("uid = ?", uid).First(&user).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}

	jwt, err := service.GenerateJWT(user)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(response.Login{
		AccessToken:  jwt,
		RefreshToken: tokenString,
		ExpiresAt:    time.Now().Add(service.AccessTokenJwtExpDuration).Format(time.RFC3339),
	})
}

func (app App) Logout(c *fiber.Ctx) error {
	req := new(request.LogoutRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	service.BlackListToken(app.cache, req.RefreshToken, time.Now().Add(12*time.Hour).Unix())
	service.BlackListToken(app.cache, c.Locals("jwt").(string), c.Locals("exp").(int64))

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

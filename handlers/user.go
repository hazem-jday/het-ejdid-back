package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hazem-jday/het-ejdid-back/config"
	"github.com/hazem-jday/het-ejdid-back/entities"
)

func Signup(c *fiber.Ctx) error {
	user := new(entities.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	config.Database.Create(&user)
	return c.Status(201).JSON(user)
}

func Login(c *fiber.Ctx) error {
	login := new(entities.Login)
	if err := c.BodyParser(login); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	username := login.Username

	var user entities.User
	result := config.Database.Find(&user, "username = ?", username)
	if result.RowsAffected == 0 {
		return c.SendStatus(403)
	}

	if login.Password == user.Password {
		return c.Status(200).JSON(&user)
	}

	return c.SendStatus(403)
}

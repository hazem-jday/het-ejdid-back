package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hazem-jday/het-ejdid-back/entities"
)

func GetTest(c *fiber.Ctx) error {
	user := entities.User{
		Username: "hazem",
		Password: "123456",
	}
	return c.Status(200).JSON(user)
}

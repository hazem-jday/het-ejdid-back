package handlers

import (
	"het-ejdid-back/entities"

	"github.com/gofiber/fiber/v2"
)

func GetTest(c *fiber.Ctx) error {
	user := entities.User{
		Username: "hazem",
		Password: "123456",
	}
	return c.Status(200).JSON(user)
}

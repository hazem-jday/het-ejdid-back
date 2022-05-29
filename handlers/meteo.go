package handlers

import (
	"het-ejdid-back/config"
	"het-ejdid-back/entities"

	"github.com/gofiber/fiber/v2"
)

func GetMeteo(c *fiber.Ctx) error {
	var meteo []entities.Meteo

	result := config.Database.Find(&meteo)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&meteo)
}

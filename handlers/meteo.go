package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hazem-jday/het-ejdid-back/config"
	"github.com/hazem-jday/het-ejdid-back/entities"
)

func GetMeteo(c *fiber.Ctx) error {
	var meteo []entities.Meteo

	result := config.Database.Find(&meteo)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&meteo)
}

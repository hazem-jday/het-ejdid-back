package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hazem-jday/het-ejdid-back/config"
	"github.com/hazem-jday/het-ejdid-back/entities"
)

func GetArticles(c *fiber.Ctx) error {
	var articles []entities.Article
	config.Database.Find(&articles)

	fmt.Print(articles)
	return c.Status(200).JSON(articles)
}

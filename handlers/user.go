package handlers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hazem-jday/het-ejdid-back/config"
	"github.com/hazem-jday/het-ejdid-back/entities"
)

func Signup(c *fiber.Ctx) error {
	user := new(entities.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	var users []entities.User
	if config.Database.Where("username = ?", user.Username).Find(&users).RowsAffected > 0 {
		return c.Status(406).JSON(user)
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

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entities.User

	result := config.Database.First(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	userUpdated := new(entities.User)
	user := new(entities.User)

	if err := c.BodyParser(userUpdated); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	config.Database.First(&user, id)

	var users []entities.User
	if strings.Compare(user.Username, userUpdated.Username) != 0 && config.Database.Where("username = ?", userUpdated.Username).Find(&users).RowsAffected > 0 {
		return c.Status(406).JSON(user)
	}

	user.FirstName = userUpdated.FirstName
	user.FamilyName = userUpdated.FamilyName
	user.Email = userUpdated.Email
	user.BirthDate = userUpdated.BirthDate
	user.Password = userUpdated.Password
	user.Username = userUpdated.Username

	config.Database.Save(&user)

	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	var user entities.User
	id := c.Params("id")
	config.Database.First(&user, id)

	config.Database.Delete(&user)

	return c.Status(200).JSON(user)
}

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hazem-jday/het-ejdid-back/config"
	"github.com/hazem-jday/het-ejdid-back/handlers"
)

func main() {
	//Initialisation du serveur
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	config.Connect()

	app.Get("/test", handlers.GetTest)

	// Démarrage
	log.Fatal(app.Listen(":8081"))
}

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

	app.Get("/articles", handlers.GetArticles)
	app.Get("/article/:id", handlers.GetArticle)

	app.Get("/interPreview", handlers.GetPreviewInter)
	app.Get("/natPreview", handlers.GetPreviewNat)
	app.Get("/sportPreview", handlers.GetPreviewSport)

	app.Get("/newsTicker", handlers.GetNewsTicker)
	app.Get("/newsHighlights", handlers.GetNewsHighlights)

	app.Get("search/:s", handlers.Search)

	app.Get("inter/:n", handlers.Inter)
	app.Get("nat/:n", handlers.Nat)
	app.Get("sport/:n", handlers.Sport)

	app.Get("nbLikes/:id", handlers.GetNBLikes)

	app.Get("like/:user/:article", handlers.GetLike)
	app.Get("save/:user/:article", handlers.GetSave)

	app.Post("/signup", handlers.Signup)
	app.Post("/login", handlers.Login)

	app.Post("/like", handlers.LikeArticle)
	app.Post("/save", handlers.SaveArticle)

	app.Delete("/unlike/:id", handlers.UnlikeArticle)
	app.Delete("/unsave/:id", handlers.UnsaveArticle)

	app.Delete("/unsave/:user/:id", handlers.UnsaveArticleReturnSaved)
	app.Delete("/unlike/:user/:id", handlers.UnlikeArticleReturnLiked)

	app.Get("saved/:user", handlers.GetSaved)
	app.Get("liked/:user", handlers.GetLiked)

	app.Get("user/:id", handlers.GetUser)

	app.Put("user/:id", handlers.UpdateUser)

	app.Delete("user/:id", handlers.DeleteUser)

	app.Get("meteo", handlers.GetMeteo)

	// Démarrage

	log.Fatal(app.Listen(":8081"))
}

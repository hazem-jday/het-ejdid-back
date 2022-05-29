package handlers

import (
	"fmt"
	"het-ejdid-back/config"
	"het-ejdid-back/entities"
	"net/url"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	var articles []entities.Article
	config.Database.Find(&articles)

	return c.Status(200).JSON(articles)
}

func AddArticle(c *fiber.Ctx) error {
	article := new(entities.Article)

	if err := c.BodyParser(article); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&article)

	return c.Status(201).JSON(article)
}
func GetPreviewInter(c *fiber.Ctx) error {
	var articles []entities.Article
	config.Database.Where("type = 'inter'").Order("date DESC").Limit(6).Find(&articles)

	return c.Status(200).JSON(articles)
}
func GetPreviewNat(c *fiber.Ctx) error {
	var articles []entities.Article
	config.Database.Where("type = 'nat'").Order("date DESC").Limit(6).Find(&articles)

	return c.Status(200).JSON(articles)
}
func GetPreviewSport(c *fiber.Ctx) error {
	var articles []entities.Article
	config.Database.Where("type = 'sport'").Order("date DESC").Limit(6).Find(&articles)

	return c.Status(200).JSON(articles)
}

func GetNewsTicker(c *fiber.Ctx) error {
	var articles []entities.Article
	config.Database.Order("date DESC").Limit(5).Find(&articles)

	return c.Status(200).JSON(articles)
}

func GetNewsHighlights(c *fiber.Ctx) error {
	var nat []entities.Article
	var inter []entities.Article
	var sport []entities.Article

	var highlights entities.NewsHighlights

	config.Database.Where("type = 'nat'").Order("date DESC").Limit(5).Find(&nat)
	config.Database.Where("type = 'inter'").Order("date DESC").Limit(5).Find(&inter)
	config.Database.Where("type = 'sport'").Order("date DESC").Limit(5).Find(&sport)

	highlights.Nat = nat
	highlights.Inter = inter
	highlights.Sport = sport

	return c.Status(200).JSON(highlights)
}

func GetArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article entities.Article

	result := config.Database.Find(&article, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&article)
}

func Search(c *fiber.Ctx) error {
	s := c.Params("s")
	search, _ := url.QueryUnescape(s)
	var articles []entities.Article
	config.Database.Where(fmt.Sprintf("%s%s%s", "convert(title using utf8) LIKE '%", search, "%'")).Order("date DESC").Limit(18).Find(&articles)
	return c.Status(200).JSON(articles)
}

func Inter(c *fiber.Ctx) error {
	n, err := strconv.Atoi(c.Params("n"))
	if err != nil || n < 1 {
		return c.SendStatus(404)
	}
	var articles []entities.Article
	config.Database.Where("type = 'inter'").Order("date DESC").Offset((n - 1) * 12).Limit(12).Find(&articles)
	return c.Status(200).JSON(articles)
}

func Nat(c *fiber.Ctx) error {
	n, err := strconv.Atoi(c.Params("n"))
	if err != nil || n < 1 {
		return c.SendStatus(404)
	}
	var articles []entities.Article
	config.Database.Where("type = 'nat'").Order("date DESC").Offset((n - 1) * 12).Limit(12).Find(&articles)
	return c.Status(200).JSON(articles)
}
func Sport(c *fiber.Ctx) error {
	n, err := strconv.Atoi(c.Params("n"))
	if err != nil || n < 1 {
		return c.SendStatus(404)
	}
	var articles []entities.Article
	config.Database.Where("type = 'sport'").Order("date DESC").Offset((n - 1) * 12).Limit(12).Find(&articles)
	return c.Status(200).JSON(articles)
}

func GetLike(c *fiber.Ctx) error {
	var like entities.Like

	userID := c.Params("user")
	articleID := c.Params("article")

	config.Database.Where("user = ? and article = ?", userID, articleID).First(&like)

	return c.Status(200).JSON(like)
}
func LikeArticle(c *fiber.Ctx) error {
	like := new(entities.Like)

	if err := c.BodyParser(like); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&like)

	return c.Status(201).JSON(like)
}
func UnlikeArticle(c *fiber.Ctx) error {
	var like entities.Like
	id := c.Params("id")

	config.Database.First(&like, id)

	if like.ID == 0 {
		fmt.Println("---------------------------------------------")
		return c.Status(200).JSON(like)
	}

	config.Database.Delete(&like)

	like.ID = 0
	like.Article = 0
	like.User = 0
	fmt.Println(like.ID)

	return c.Status(200).JSON(like)
}

func GetNBLikes(c *fiber.Ctx) error {
	var n int64
	var nb entities.NBLikes

	articleID := c.Params("id")

	config.Database.Model(&entities.Like{}).Where("article = ?", articleID).Count(&n)
	nb.NB = uint(n)
	idInt, _ := strconv.Atoi(articleID)

	nb.ID = uint(idInt)

	return c.Status(200).JSON(nb)
}

func GetSave(c *fiber.Ctx) error {
	var save entities.Save

	userID := c.Params("user")
	articleID := c.Params("article")

	config.Database.Where("user = ? and article = ?", userID, articleID).First(&save)

	return c.Status(200).JSON(save)
}
func SaveArticle(c *fiber.Ctx) error {
	save := new(entities.Save)

	if err := c.BodyParser(save); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&save)

	return c.Status(201).JSON(save)
}
func UnsaveArticle(c *fiber.Ctx) error {
	var save entities.Save
	id := c.Params("id")

	config.Database.First(&save, id)

	if save.ID == 0 {
		fmt.Println("---------------------------------------------")
		return c.Status(200).JSON(save)
	}

	config.Database.Delete(&save)

	save.ID = 0
	save.Article = 0
	save.User = 0
	fmt.Println(save.ID)

	return c.Status(200).JSON(save)
}

func UnsaveArticleReturnSaved(c *fiber.Ctx) error {
	var save entities.Save
	var saves []entities.Save
	id := c.Params("id")
	user := c.Params("user")
	config.Database.First(&save, id)

	if save.ID == 0 {
		fmt.Println("---------------------------------------------")
		return c.Status(200).JSON(save)
	}

	config.Database.Delete(&save)

	config.Database.Where("user = ?", user).Find(&saves)

	return c.Status(200).JSON(saves)
}

func GetSaved(c *fiber.Ctx) error {
	var saves []entities.Save
	userID := c.Params("user")

	config.Database.Where("user = ?", userID).Find(&saves)

	return c.Status(200).JSON(saves)
}

func GetLiked(c *fiber.Ctx) error {
	var likes []entities.Like
	userID := c.Params("user")

	config.Database.Where("user = ?", userID).Find(&likes)

	return c.Status(200).JSON(likes)
}

func UnlikeArticleReturnLiked(c *fiber.Ctx) error {
	var like entities.Like
	var likes []entities.Like
	id := c.Params("id")
	user := c.Params("user")
	config.Database.First(&like, id)

	if like.ID == 0 {
		fmt.Println("---------------------------------------------")
		return c.Status(200).JSON(like)
	}

	config.Database.Delete(&like)

	config.Database.Where("user = ?", user).Find(&likes)

	return c.Status(200).JSON(likes)
}

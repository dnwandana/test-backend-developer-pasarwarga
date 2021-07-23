package controller

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/service"
	"github.com/gofiber/fiber/v2"
)

type ArticleController struct {
	articleService *service.ArticleService
}

func NewArticleController(articleService *service.ArticleService) ArticleController {
	return ArticleController{
		articleService: articleService,
	}
}

func (controller *ArticleController) SetupRoutes(app *fiber.App) {
	app.Post("/article", controller.Create)
	app.Get("/article", controller.List)
	app.Get("/article/:id", controller.FindOne)
	app.Put("/article/:id", controller.Update)
	app.Delete("/article/:id", controller.SoftDelete)
	app.Get("/article/deleted", controller.ListSoftDeleted)
	app.Delete("/article/deleted/:id", controller.Delete)
}

func (controller *ArticleController) Create(ctx *fiber.Ctx) error {
	return ctx.SendString("Create new article")
}

func (controller *ArticleController) List(ctx *fiber.Ctx) error {
	return ctx.SendString("List all article")
}

func (controller *ArticleController) FindOne(ctx *fiber.Ctx) error {
	return ctx.SendString("Find specific article")
}

func (controller *ArticleController) Update(ctx *fiber.Ctx) error {
	return ctx.SendString("Update article")
}

func (controller *ArticleController) SoftDelete(ctx *fiber.Ctx) error {
	return ctx.SendString("Soft delete article")
}

func (controller *ArticleController) ListSoftDeleted(ctx *fiber.Ctx) error {
	return ctx.SendString("List all soft deleted article")
}

func (controller *ArticleController) Delete(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete article")
}

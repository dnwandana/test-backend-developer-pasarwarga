package controller

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/service"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	categoryService *service.CategoryService
}

func NewCategoryController(categoryService *service.CategoryService) CategoryController {
	return CategoryController{
		categoryService: categoryService,
	}
}

func (controller *CategoryController) SetupRoutes(app *fiber.App) {
	app.Post("/category", controller.Create)
	app.Get("/category", controller.List)
	app.Get("/category/:id", controller.FindOne)
	app.Put("/category/:id", controller.Update)
	app.Delete("/category/:id", controller.SoftDelete)
	app.Get("/category/deleted", controller.ListSoftDeleted)
	app.Delete("/category/deleted/:id", controller.Delete)
}

func (controller *CategoryController) Create(ctx *fiber.Ctx) error {
	return ctx.SendString("Create new category")
}

func (controller *CategoryController) List(ctx *fiber.Ctx) error {
	return ctx.SendString("List all category")
}

func (controller *CategoryController) FindOne(ctx *fiber.Ctx) error {
	return ctx.SendString("Find specific category")
}

func (controller *CategoryController) Update(ctx *fiber.Ctx) error {
	return ctx.SendString("Update specific category")
}

func (controller *CategoryController) SoftDelete(ctx *fiber.Ctx) error {
	return ctx.SendString("Soft delete category")
}

func (controller *CategoryController) ListSoftDeleted(ctx *fiber.Ctx) error {
	return ctx.SendString("List all soft deleted category")
}

func (controller *CategoryController) Delete(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete category")
}

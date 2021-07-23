package main

import (
	"database/sql"
	"fmt"
	"github.com/dnwandana/test-backend-developer-pasarwarga/controller"
	"github.com/dnwandana/test-backend-developer-pasarwarga/model"
	"github.com/dnwandana/test-backend-developer-pasarwarga/repository"
	"github.com/dnwandana/test-backend-developer-pasarwarga/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"time"
)

var (
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
)

func main() {
	// get connection
	dbUri := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	connection, err := sql.Open("mysql", dbUri)
	if err != nil {
		panic(err)
	}

	err = connection.Ping()
	if err != nil {
		panic(err)
	}

	// set connection pool
	connection.SetMaxIdleConns(5)
	connection.SetMaxOpenConns(100)
	connection.SetConnMaxIdleTime(5 * time.Minute)
	connection.SetConnMaxLifetime(60 * time.Minute)

	// setup article repository, service, and controller
	articleRepository := repository.NewArticleRepository(connection)
	articleService := service.NewArticleService(&articleRepository)
	articleController := controller.NewArticleController(&articleService)

	// setup category repository, service, and controller
	categoryRepository := repository.NewCategoryRepository(connection)
	categoryService := service.NewCategoryService(&categoryRepository)
	categoryController := controller.NewCategoryController(&categoryService)

	// new fiber app
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
					StatusCode: fiber.StatusBadRequest,
					Error:      err.Error(),
				})
			}
			return nil
		},
	})

	// cors middleware
	app.Use(cors.New())

	// receiver custom error handler
	app.Use(recover.New())

	// setup routes
	articleController.SetupRoutes(app)
	categoryController.SetupRoutes(app)

	log.Fatal(app.Listen(":5000"))
}

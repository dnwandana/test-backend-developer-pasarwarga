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
	Connection *sql.DB
)

// initialize database connection
func init() {
	var err error
	dbUri := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	// open connection
	Connection, err = sql.Open("mysql", dbUri)
	if err != nil {
		panic(err)
	}

	// check connection
	err = Connection.Ping()
	if err != nil {
		panic(err)
	}

	// set connection pool
	Connection.SetMaxIdleConns(5)
	Connection.SetMaxOpenConns(100)
	Connection.SetConnMaxIdleTime(5 * time.Minute)
	Connection.SetConnMaxLifetime(60 * time.Minute)
}

func main() {
	// setup article repository, service, and controller
	articleRepository := repository.NewArticleRepository(Connection)
	articleService := service.NewArticleService(&articleRepository)
	articleController := controller.NewArticleController(&articleService)

	// setup category repository, service, and controller
	categoryRepository := repository.NewCategoryRepository(Connection)
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

	// listen to port 5000
	log.Fatal(app.Listen(":5000"))
}

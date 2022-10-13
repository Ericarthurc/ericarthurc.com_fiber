package main

import (
	"ericarthurc/router"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django"
	"github.com/joho/godotenv"
)

func main() {
	// Load environmental variables
	godotenv.Load()

	// Set template engine
	engine := django.New("./views", ".django")

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		Views:         engine,
		ErrorHandler:  errorHandler,
	})

	// Temp error handling
	app.Use(recover.New())

	// Serve static routes
	app.Static("/public", "./public")
	app.Static("/sw.js", "./public/js/sw.js")
	app.Static("/robots.txt", "./public/robots.txt")

	// Serve router
	router.SetupRoutes(app)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	return c.Status(code).Render("error", fiber.Map{
		"error": err.Error(),
		"code":  code,
		"uri":   c.Request().URI(),
	})
}

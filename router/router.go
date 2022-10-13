package router

import (
	"ericarthurc/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	blog := app.Group("/")
	blog.Get("/", handlers.BlogIndex)
	blog.Get("/blog/:blog", handlers.GetBlog)

	series := app.Group("/series")
	series.Get("/", handlers.SeriesIndex)
	series.Get("/:series", handlers.GetSeries)

	category := app.Group("/category")
	category.Get("/:category", handlers.GetCategory)

	about := app.Group("/about")
	about.Get("/", handlers.AboutIndex)

	// admin := app.Group("/admin")
	// admin.Get("/")

}

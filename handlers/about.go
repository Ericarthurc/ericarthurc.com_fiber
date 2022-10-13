package handlers

import "github.com/gofiber/fiber/v2"

func AboutIndex(c *fiber.Ctx) error {
	return c.Render("about", nil)
}

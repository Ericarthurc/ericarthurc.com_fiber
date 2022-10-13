package handlers

import (
	"ericarthurc/utils"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GetCategory(c *fiber.Ctx) error {
	category := c.Params("category")

	categoryUnescaped, err := url.QueryUnescape(category)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	categoryList, err := utils.GetSliceByCategoryParam(categoryUnescaped)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Render("category", fiber.Map{"category": categoryUnescaped, "categoryList": categoryList})
}

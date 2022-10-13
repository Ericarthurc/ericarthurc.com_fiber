package handlers

import (
	"ericarthurc/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func BlogIndex(c *fiber.Ctx) error {

	posts, err := utils.GetBlogPostsSlice()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Render("index", fiber.Map{"blogList": posts})
}

func GetBlog(c *fiber.Ctx) error {
	blogName := c.Params("blog")

	markdown, meta, err := utils.GetBlogPostByName(blogName)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("blog post %s was not found", blogName))
	}

	return c.Render("blog", fiber.Map{"markdown": markdown, "meta": meta})
}

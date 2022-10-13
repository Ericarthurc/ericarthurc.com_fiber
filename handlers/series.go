package handlers

import (
	"ericarthurc/utils"

	"github.com/gofiber/fiber/v2"
)

func SeriesIndex(c *fiber.Ctx) error {
	series, err := utils.GetSeriesSlice()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Render("index_series", fiber.Map{"seriesList": series})
}

func GetSeries(c *fiber.Ctx) error {
	series := c.Params("series")

	seriesList, err := utils.GetSliceBySeriesParam(series)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Render("series", fiber.Map{"series": series, "seriesList": seriesList})
}

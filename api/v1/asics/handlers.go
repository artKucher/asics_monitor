package asics

import "github.com/gofiber/fiber/v2"

func AsicsList(c *fiber.Ctx) error {
	asics := []string{"asics1", "asics2", "asics3"}

	return c.JSON(fiber.Map{
		"success": true,
		"asics":   asics,
	})
}

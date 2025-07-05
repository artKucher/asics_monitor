package statistics

import "github.com/gofiber/fiber/v2"

func GetStatistics(c *fiber.Ctx) error {
	statistics := []string{"1", "2", "3"}

	return c.JSON(fiber.Map{
		"success":    true,
		"statistics": statistics,
	})
}

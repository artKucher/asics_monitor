package main

import (
	asicsHandlers "asics_monitor/api/v1/asics"
	statisticsHandlers "asics_monitor/api/v1/statistics"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := Setup()
	app.Listen("0.0.0.0:3000")
}

func Setup() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		GETOnly:           false, /* TODO: change to true */
		ServerHeader:      "Asics Monitor",
		AppName:           "Asics Monitor v0.0.1",
	})
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	v1 := app.Group("api/v1/")

	v1.Get("/asics/", asicsHandlers.AsicsList)
	v1.Get("/statistics/", statisticsHandlers.GetStatistics)
	return app
}

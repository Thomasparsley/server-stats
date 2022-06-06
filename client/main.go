package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

const (
	PORT = ":3000"
)

func main() {
	//s, _ := NewStatistic()

	app := fiber.New(fiber.Config{
		ServerHeader: "vel",
	})

	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(compress.New())
	app.Use(helmet.New())

	app.Get("/stats", func(c *fiber.Ctx) error {
		return nil
	})

	app.Listen(PORT)
}

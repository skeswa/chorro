package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/skeswa/chorro/apps/server/config"
)

func main() {
	configuration, err := config.New()
	if err == nil {
		fmt.Printf("Configuration:\n%+v\n", configuration)
	} else {
		log.Fatalln("Failed to read configuration:", err)
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, API 👋!")
	})

	app.Listen(fmt.Sprintf(":%d", configuration.HttpPort))
}

package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/skeswa/chorro/apps/server/config"
)

func main() {
	c, err := config.ReadConfig()
	if err != nil {
		log.Fatalln("Failed to read configuration", err)
	}

	app := fiber.New()

	if c.Environment == config.DevEnvironment {
		fmt.Println("dev modddeeee")
		app.Use(cors.New(cors.Config{
			// Allow requests from the web frontend locally.
			AllowOrigins: "http://localhost:3000",
		}))
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(fmt.Sprintf(":%d", c.HttpPort))
}

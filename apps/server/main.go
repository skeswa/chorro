package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/shareed2k/goth_fiber"
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

	goth.UseProviders(
		google.New(configuration.Auth.GoogleOAuth2ClientID, configuration.Auth.GoogleOAuth2ClientSecret, fmt.Sprintf("%s/auth/google/callback", configuration.PublicBaseUrl()), "email", "profile"),
	)

	app.Get("/login/:provider", goth_fiber.BeginAuthHandler)
	app.Get("/auth/:provider/callback", func(ctx *fiber.Ctx) error {
		user, err := goth_fiber.CompleteUserAuth(ctx)
		if err != nil {
			return ctx.SendString("wtf")
		}

		return ctx.SendString(fmt.Sprintf("%+v", user))
	})

	app.Listen(fmt.Sprintf(":%d", configuration.HttpPort))
}

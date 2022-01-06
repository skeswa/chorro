package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/skeswa/chorro/apps/server/config"
	"github.com/skeswa/goth_fiber"
)

func main() {
	configuration, err := config.New()
	if err == nil {
		fmt.Printf("Configuration:\n%+v\n", configuration)
	} else {
		log.Fatalln("Failed to read configuration:", err)
	}

	// For session management, we talk to Redis.
	fiberRedisStore := redis.New(configuration.ForFiberRedis())
	sessionStore := session.New(configuration.ForFiberSession(fiberRedisStore))

	// We use Goth to manage authentication against 3rd party providers. To add
	// compatibility between our web server, Fiber, and Goth, we use a cool little
	// library called goth_fiber. It needs to be told about our session store so
	// it can automate OAuth handshake.
	goth_fiber.SessionStore = sessionStore
	goth.UseProviders(
		// Add support for "Login With Google"
		google.New(
			configuration.Auth.GoogleOAuth2ClientID,
			configuration.Auth.GoogleOAuth2ClientSecret,
			// Fully-qualified Google auth callback endpoint URL:
			fmt.Sprintf("%s/auth/google/callback", configuration.PublicBaseUrl()),
			// Google OAuth 2.0 scopes:
			"email", "profile",
		),
	)

	app := fiber.New()

	// Enable incoming requests from different domains. This is important for
	// local development especially since the web app and API server are served
	// from different ports on `localhost`. In production, CORS support simply
	// gives us the flexibility to host API on a different subdomain, for example.
	app.Use(cors.New())

	// Route used to test that the server is up. This will likely disappear in
	// future.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// Future placeholder for the GraphQL API.
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, API 👋!")
	})

	// Used to check whether the user is logged in.
	app.Get("/me", func(c *fiber.Ctx) error {
		// Read the session from the store.
		session, err := sessionStore.Get(c)
		if err != nil {
			log.Println("Failed to read the session:", err)

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		log.Println("session1", session.ID(), session.Keys())

		// If the user is already logged, there is no need to complete user auth.
		userEmail, isUserLoggedIn := session.Get("email").(string)
		if isUserLoggedIn {
			return c.SendString(userEmail)
		}

		return c.SendStatus(fiber.StatusUnauthorized)
	})

	// Starts the Goth login flow.
	app.Get("/login/:provider", func(c *fiber.Ctx) error {
		// Read the session from the store.
		session, err := sessionStore.Get(c)
		if err != nil {
			log.Println("Failed to read the session:", err)

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		log.Println("session2", session.ID(), session.Keys())

		// If the user is already logged, there is no need to complete user auth.
		userEmail, isUserLoggedIn := session.Get("email").(string)
		if isUserLoggedIn {
			return c.SendString(fmt.Sprintf("Hello, %s! No need to login again!", userEmail))
		}

		return goth_fiber.BeginAuthHandler(c)
	})

	// Finishes the Goth login flow.
	app.Get("/auth/:provider/callback", func(c *fiber.Ctx) error {
		// Read the session from the store.
		session, err := sessionStore.Get(c)
		if err != nil {
			log.Println("Failed to read the session:", err)

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		log.Println("session3", session.ID(), session.Keys())

		// If the user is already logged, there is no need to complete user auth.
		userEmail, isUserLoggedIn := session.Get("email").(string)
		if isUserLoggedIn {
			return c.SendString(fmt.Sprintf("Hello, %s! You're already logged in!", userEmail))
		}

		// If the user is not logged in, we need to finish authenticating.
		gothUser, err := goth_fiber.CompleteUserAuth(c, goth_fiber.CompleteUserAuthOptions{
			ShouldLogout: false,
		})
		if err != nil {
			log.Println("Failed to complete user auth:", err)

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		userEmail = gothUser.Email
		session.Set("email", userEmail)
		if err = session.Save(); err != nil {
			log.Println("Failed to save the user's email to the session:", err)

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendString(fmt.Sprintf("Hello, %s! You just logged in!", userEmail))
	})

	app.Listen(fmt.Sprintf(":%d", configuration.HttpPort))
}

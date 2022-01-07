package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/skeswa/chorro/apps/server/config"
	"github.com/skeswa/chorro/apps/server/db"
	"github.com/skeswa/chorro/apps/server/session"
	"github.com/skeswa/goth_fiber"
)

func main() {
	// Read configuration from the environment and arguments.
	configuration, err := config.New()
	if err != nil {
		log.Fatalln("Failed to read configuration:", err)
	} else {
		fmt.Printf("Configuration:\n%+v\n", configuration)
	}

	// Connect to the database.
	db, err := db.New(configuration)
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
	}

	// For session management, we talk to Redis.
	sessionStore := session.New(configuration)

	app := fiber.New()

	// Enable incoming requests from different domains. This is important for
	// local development especially since the web app and API server are served
	// from different ports on `localhost`. In production, CORS support simply
	// gives us the flexibility to host API on a different subdomain, for example.
	app.Use(cors.New())

	// Setup Google 3rd-party auth via the Goth authentication library.
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
	sessionStore.BindToGoth()

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
		session, err := sessionStore.Session(c)
		if err != nil {
			log.Println("Failed to read the session:", err)

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if !session.IsUserLoggedIn {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		user, err := db.User(session.UserID)
		if err != nil {
			log.Println("/me failed:", err)

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if user == nil {
			if err := session.LogOut(); err != nil {
				log.Printf("Failed to log out unlisted user \"%d\"\n", session.UserID)

				return c.SendStatus(fiber.StatusInternalServerError)
			} else {
				fmt.Printf("Logged out unlisted user \"%d\"\n", session.UserID)

				return c.SendStatus(fiber.StatusNotFound)
			}
		}

		return c.SendString(fmt.Sprintf("%+v", user))
	})

	// Starts the "Login With Google" flow.
	app.Get("/login/:provider", func(c *fiber.Ctx) error {
		session, err := sessionStore.Session(c)
		if err != nil {
			log.Println("Failed to read the session:", err)

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Go back home if the user is already logged in.
		if session.IsUserLoggedIn {
			if doesUserExist, err := db.DoesUserExist(session.UserID); doesUserExist && err == nil {
				return c.Redirect(configuration.HomeUrl)
			} else if !doesUserExist {
				log.Println("Tried to short-circuit login, but user did not exist")
			} else {
				log.Println(
					"Tried to short-circuit login, but there was an error:",
					err,
				)
			}
		}

		// Kick off "Login With Google" flow.
		return goth_fiber.BeginAuthHandler(c)
	})

	homeUrlWithAuthFailureFlag := fmt.Sprintf("%s?auth=failed", configuration.HomeUrl)

	// Finishes the "Login With Google".
	app.Get("/auth/:provider/callback", func(c *fiber.Ctx) error {
		session, err := sessionStore.Session(c)
		if err != nil {
			log.Println("Failed to read the session:", err)

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Go back home if the user is already logged in.
		if session.IsUserLoggedIn {
			if doesUserExist, err := db.DoesUserExist(session.UserID); doesUserExist && err == nil {
				return c.Redirect(configuration.HomeUrl)
			} else if !doesUserExist {
				log.Println("Tried to short-circuit login, but user did not exist")
			} else {
				log.Println(
					"Tried to short-circuit login, but there was an error:",
					err,
				)
			}
		}

		googleUser, err := goth_fiber.CompleteUserAuth(
			c, goth_fiber.CompleteUserAuthOptions{
				ShouldLogout: false,
			},
		)
		if err != nil {
			log.Println("Failed to complete \"Login With Google\":", err)

			return c.Redirect(homeUrlWithAuthFailureFlag)
		}

		if userIDOfGoogleUser, err := db.UserIDOfGoogleUser(
			googleUser.UserID,
		); err != nil {
			log.Println("Failed to complete \"Login With Google\":", err)

			return c.Redirect(homeUrlWithAuthFailureFlag)
		} else {
			if userIDOfGoogleUser == nil {
				// This is an unrecognized Google user, so we should register them.
				if userID, err := db.RegisterGoogleUser(googleUser); err != nil {
					log.Println("Failed to register Google user:", err)

					return c.Redirect(homeUrlWithAuthFailureFlag)
				} else {
					userIDOfGoogleUser = &userID
				}
			}

			if err := session.LogIn(*userIDOfGoogleUser); err != nil {
				log.Println("Failed to log in recognized Google user:", err)

				return c.Redirect(homeUrlWithAuthFailureFlag)
			}
		}

		// Looks like the user is now logged in - let's send them home.
		return c.Redirect(configuration.HomeUrl)
	})

	app.Listen(fmt.Sprintf(":%d", configuration.HttpPort))
}

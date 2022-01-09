package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/rs/cors"
	"github.com/skeswa/chorro/apps/server/auth"
	"github.com/skeswa/chorro/apps/server/cache"
	"github.com/skeswa/chorro/apps/server/config"
	"github.com/skeswa/chorro/apps/server/db"
	"github.com/skeswa/chorro/apps/server/graph"
	"github.com/skeswa/chorro/apps/server/session"
)

func main() {
	// Set the number of CPUs that can execute concurrently to the number cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Read configuration from the environment and arguments.
	configuration, err := config.New()
	if err != nil {
		log.Fatalln("Failed to read configuration:", err)
	} else {
		fmt.Printf("Configuration:\n%+v\n", configuration)
	}

	// Connect to the cache.
	cache, err := cache.New(&configuration)
	if err != nil {
		log.Fatalln("Failed to connect to the cache:", err)
	}

	// Connect to the database.
	db, err := db.New(&configuration)
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
	}

	// Create a new, non-default HTTP serving multiplexer so we can mess with it.
	mux := http.NewServeMux()

	// Hook up authentication.
	auth.Setup(cache, &configuration, db, mux)

	// Hook up the GraphQL API.
	graph.Setup(mux)

	// Routed used to test that the server is up. This will likely disappear in
	// future.
	mux.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		// GET "/"" only - everything else is a 404.
		if request.Method != http.MethodGet || request.URL.Path != "/" {
			responseWriter.WriteHeader(http.StatusNotFound)

			return
		}

		fmt.Fprintf(responseWriter, "Hello, World 👋!\n")
	})

	// Useful debugging route that shits out the authenticated user.
	mux.HandleFunc("/me", func(responseWriter http.ResponseWriter, request *http.Request) {
		// GET only.
		if request.Method != http.MethodGet {
			responseWriter.WriteHeader(http.StatusNotFound)

			return
		}

		session := session.Read(cache, request, responseWriter)
		if !session.IsUserLoggedIn {
			responseWriter.WriteHeader(http.StatusUnauthorized)

			return
		}

		user, err := db.User(session.UserID)
		if err != nil {
			log.Println("/me failed:", err)

			responseWriter.WriteHeader(http.StatusInternalServerError)

			return
		}

		fmt.Fprintf(responseWriter, "%+v", user)
	})

	// Setup CORS middleware.
	cors := cors.New(configuration.ForCors())
	muxWithCors := cors.Handler(mux)

	// Arrite! Time to start 'er up.
	fmt.Println()
	log.Printf("HTTP server listening on port %d...\n", configuration.HttpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", configuration.HttpPort), muxWithCors); err != nil {
		log.Fatalln("Failed to start the server:", err)
	}
}

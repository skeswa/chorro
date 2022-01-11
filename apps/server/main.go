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
	"github.com/skeswa/chorro/apps/server/mailer"
	"github.com/skeswa/chorro/apps/server/session"
)

func main() {
	// Set the number of CPUs that can execute concurrently to the number cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Read configuration from the environment and arguments.
	config, err := config.New()
	if err != nil {
		log.Fatalln("Failed to read configuration:", err)
	} else {
		fmt.Printf("Config:\n%+v\n", config)
	}

	// Connect to the cache.
	cache, err := cache.New(config)
	if err != nil {
		log.Fatalln("Failed to connect to the cache:", err)
	}

	// Connect to the database.
	db, err := db.New(config)
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
	}

	// Connect to the mail sending service.
	mailer := mailer.New(config)

	// Create a new, non-default HTTP serving multiplexer so we can mess with it.
	mux := http.NewServeMux()

	// Hook up authentication.
	auth.Setup(cache, config, db, mux)

	// Hook up the GraphQL API.
	graph.Setup(cache, config, db, mailer, mux)

	// Setup CORS middleware.
	cors := cors.New(config.ForCors())

	// Bind middleware to mux before starting our engines.
	muxWithMiddleware := cors.Handler(mux)
	muxWithMiddleware = session.Handler(cache, muxWithMiddleware)

	// Arrite! Time to start 'er up.
	fmt.Println()
	log.Printf("HTTP server listening on port %d...\n", config.HttpPort)
	if err := http.ListenAndServe(
		fmt.Sprintf(":%d", config.HttpPort),
		muxWithMiddleware,
	); err != nil {
		log.Fatalln("Failed to start the server:", err)
	}
}

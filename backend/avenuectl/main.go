package main

import (
	"avenue/backend/handlers"
	"avenue/backend/persist"
)

func main() {
	persist := persist.NewPersist(
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "user"),
		getEnv("DB_PASSWORD", "secret"),
		getEnv("DB_DATABASE", "avenue"),
	)

	server := handlers.SetupServer(persist)

	server.SetupRoutes()

	// Start the server
	_ = server.Run(":8080")
}

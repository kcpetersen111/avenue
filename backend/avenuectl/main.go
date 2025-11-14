package main

import "avenue/backend/handlers"

func main() {

	server := handlers.SetupServer()

	server.SetupRoutes()

	// Start the server
	server.Run(":8080")
}

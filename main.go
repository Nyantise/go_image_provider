package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nyantise/image_provider/internal/handlers"
	"github.com/Nyantise/image_provider/internal/routes"
	"github.com/Nyantise/image_provider/serverconfig"
)

func main() {
	// Load Config
	config, err := serverconfig.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}
	// Connect to DB
	// db := dbconfig.ConnectDB(config.DatabaseURL)
	// defer db.Close()

	// Setup Handler
	handler := handlers.NewHandlers()
	// Setup HTTP server
	mux := http.NewServeMux()
	// Setup Routes
	routes.SetupRoutes(mux, handler)

	// Server instance
	serverAddr := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	fmt.Printf("Server up and listening on PORT: [%+v]\n", config.ServerPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed %v", err)
	}
}

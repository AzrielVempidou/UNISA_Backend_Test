package main

import (
	"log"
	"net/http"
	"os"

	"UNISA_Server/config"
	"UNISA_Server/routes"
	"UNISA_Server/seed"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get MongoDB URI from environment variables
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI not set in .env file")
	}

	// // Connect to MongoDB
	config.Connect(mongoURI)
	database := config.MongoClient.Database("mydatabase")

	// Run seeders
	if err := seed.SeedListTautan(database); err != nil {
		log.Fatalf("Error seeding list_tautan: %v", err)
	}
	if err := seed.SeedDataLeads(database); err != nil {
		log.Fatalf("Error seeding data_leads: %v", err)
	}

	// Create qrcodes directory if it doesn't exist
	if err := os.MkdirAll("qrcodes", os.ModePerm); err != nil {
		log.Fatalf("Failed to create qrcodes directory: %v", err)
	}

	// Create router
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Get port from environment or use default port 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

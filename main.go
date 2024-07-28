package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"www.rsvpme.com/initializers"
	"www.rsvpme.com/routes"
)

func initDB() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Use MySQL domain and make the connection
	initializers.ConnectDB()
}

func main() {
	// Initialize the database
	initDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8085")
}

// {
//     "name": "John Doe",
//     "email": "johndoe@example.com",
//     "persons": 3
// }

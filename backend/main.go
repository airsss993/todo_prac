package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"trackit/backend/initializers"
	"trackit/backend/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	initializers.ConnToDB()

	r := gin.Default()
	routes.UseRoutes(r)

	log.Println("Server started on http://localhost:8080")
	err := r.Run(":8080")
	if err != nil {
		log.Println("Error starting server")
		return
	}
}

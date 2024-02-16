package main

import (
	"log"
	"os"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}

	port := os.Getenv("PORT")

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(port)

}

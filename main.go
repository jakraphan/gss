package main

import (
	"go-rest-api/config"
	"go-rest-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	defer config.CloseDB()

	r := gin.Default()
	routes.Serve(r)
	r.Run()
}

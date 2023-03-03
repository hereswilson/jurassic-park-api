package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/database"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	r := gin.Default()

	database.Connect()

	r.Run()
}

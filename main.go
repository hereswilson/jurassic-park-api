package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/database"
	"github.com/hereswilson/jurassic-park-api/routes"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	database.MigrateDB(db)
	database.SeedDB(db)
}

func main() {
	loadEnv()
	r := gin.Default()

	router := r.Group("/api/v1")
	routes.AddRoutes(router)

	loadDatabase()

	r.Run()
}

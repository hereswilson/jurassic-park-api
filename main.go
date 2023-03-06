package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hereswilson/jurassic-park-api/database"
	"github.com/hereswilson/jurassic-park-api/routes"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() (db *gorm.DB) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	database.MigrateDB(db)
	database.SeedDB(db)
	return db
}

func main() {
	loadEnv()
	r := gin.Default()
	db := loadDatabase()
	router := r.Group("/api/v1")
	routes.AddRoutes(router, db)

	r.Run()
}

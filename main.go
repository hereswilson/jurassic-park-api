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

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading environment file")
	}
	err = godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env environment file")
	}
	err = godotenv.Load(".ENV")
	if err != nil {
		log.Println("Error loading .ENV environment file")
	}
	err = godotenv.Load("ENV")
	if err != nil {
		log.Println("Error loading ENV environment file")
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

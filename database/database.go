package database

import (
	"fmt"
	"os"

	"github.com/hereswilson/jurassic-park-api/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	var host, username, password, databaseName, port string
	_, local := os.LookupEnv("DB_HOST")
	if local {
		host = os.Getenv("DB_HOST")
		username = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		databaseName = os.Getenv("DB_NAME")
		port = os.Getenv("DB_PORT")
	} else {
		viper.SetConfigFile("ENV")
		viper.ReadInConfig()
		viper.AutomaticEnv()

		host = fmt.Sprint(viper.Get("DB_HOST"))
		username = fmt.Sprint(viper.Get("DB_USER"))
		password = fmt.Sprint(viper.Get("DB_PASSWORD"))
		databaseName = fmt.Sprint(viper.Get("DB_NAME"))
		port = fmt.Sprint(viper.Get("DB_PORT"))
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	return db, nil
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Species{}, &models.Dinosaur{}, &models.Cage{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}
	return nil
}

func SeedDB(db *gorm.DB) error {
	// Seed the database with initial data
	brachiosaurus := &models.Species{Species: "Brachiosaurus", Diet: "Herbivore"}
	stegosaurus := &models.Species{Species: "Stegosaurus", Diet: "Herbivore"}
	ankylosaurus := &models.Species{Species: "Ankylosaurus", Diet: "Herbivore"}
	triceratops := &models.Species{Species: "Triceratops", Diet: "Herbivore"}
	tyrannosaurus := &models.Species{Species: "Tyrannosaurus", Diet: "Carnivore"}
	velociraptor := &models.Species{Species: "Velociraptor", Diet: "Carnivore"}
	spinosaurus := &models.Species{Species: "Spinosaurus", Diet: "Carnivore"}
	megalosaurus := &models.Species{Species: "Megalosaurus", Diet: "Carnivore"}
	err := db.Create(brachiosaurus).Error
	if err != nil {
		return fmt.Errorf("failed to seed database: %v", err)
	}
	err = db.Create(stegosaurus).Error
	if err != nil {
		return fmt.Errorf("failed to seed database: %v", err)
	}
	err = db.Create(ankylosaurus).Error
	if err != nil {
		return fmt.Errorf("failed to seed database: %v", err)
	}
	err = db.Create(triceratops).Error
	if err != nil {
		return fmt.Errorf("failed to seed database: %v", err)
	}
	err = db.Create(tyrannosaurus).Error
	if err != nil {
		return fmt.Errorf("failed to seed database: %v", err)
	}
	err = db.Create(velociraptor).Error
	if err != nil {
		return fmt.Errorf("failed to seed database: %v", err)
	}
	err = db.Create(spinosaurus).Error
	if err != nil {
		return fmt.Errorf("failed to seed database: %v", err)
	}
	err = db.Create(megalosaurus).Error
	if err != nil {
		return fmt.Errorf("failed to seed database: %v", err)
	}

	return nil
}

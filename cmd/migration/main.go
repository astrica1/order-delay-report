package main

import (
	"fmt"
	"log"
	"os"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost     string
	dbPort     string
	dbName     string
	dbUserName string
	dbPassword string
)

func Connect2DB() *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUserName, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to postgres: %v\n", err)
	}

	return db
}

func MigrateAll(db *gorm.DB, modelTypes []interface{}) error {
	for _, model := range modelTypes {
		if err := db.AutoMigrate(&model); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	log.SetPrefix("Migration | ")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUserName = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
}

func main() {
	log.Println("Connecting to database...")
	db := Connect2DB()
	log.Println("connected to database successfully!")

	log.Println("Starting Migration...")
	modelTypes := []interface{}{
		models.Agent{},
		models.Courier{},
		models.Customer{},
		models.Vendor{},
		models.Order{},
		models.Trip{},
		models.DelayReport{},
	}

	err := MigrateAll(db, modelTypes)
	if err != nil {
		log.Fatalf("error migrating models: %v\n", err)
	}

	log.Println("models migrated successfully!")
}

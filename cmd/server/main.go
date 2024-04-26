package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/astrica1/order-delay-report/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	srvPort       string
	dbHost        string
	dbPort        string
	dbName        string
	dbUserName    string
	dbPassword    string
	redisAddress  string
	redisPassword string
	redisDB       int
)

func init() {
	log.SetPrefix("Server | ")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	srvPort = os.Getenv("SRV_PORT")
	if srvPort == "" {
		srvPort = "8080"
	}

	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUserName = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")

	redisAddress = os.Getenv("REDIS_ADDR")
	redisPassword = os.Getenv("REDIS_PASSWORD")
	redisDB, err = strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		redisDB = 0
	}
}

func Connect2DB() *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUserName, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to postgres: %v\n", err)
	}

	return db
}

func Connect2Redis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       redisDB,
	})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	log.Println("Connecting to database...")
	db := Connect2DB()
	log.Println("connected to database successfully!")

	router := gin.Default()
	router.GET("/ping", ping)

	routers.SetupAgentRouter(router, db)
	routers.SetupOrderRouter(router, db)

	router.Run(":" + srvPort)
}

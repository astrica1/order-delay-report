package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/repositories"
	"github.com/icrowley/fake"
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

func fakePhoneNumber() string {
	prefix := []string{"912", "919", "922", "935", "938", "998", "999"}
	rand.Seed(time.Now().UnixNano())
	chosenPrefix := prefix[rand.Intn(len(prefix))]
	randomDigits := fmt.Sprintf("%07d", rand.Intn(10000000))
	phoneNumber := chosenPrefix + randomDigits
	return phoneNumber
}

func Connect2DB() *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUserName, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to postgres: %v\n", err)
	}

	return db
}

func init() {
	log.SetPrefix("Seeder | ")

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

func addAgents(db *gorm.DB) error {
	r := repositories.NewAgentRepository(db)
	for i := 0; i < 20; i++ {
		m := &models.Agent{
			Username: fake.UserName(),
			Password: fake.Password(32, 32, true, true, false),
			IsActive: true,
		}

		if err := r.Create(context.TODO(), m); err != nil {
			return err
		}
	}

	return nil
}

func addCourier(db *gorm.DB) error {
	r := repositories.NewCourierRepository(db)
	for i := 0; i < 20; i++ {
		m := &models.Courier{
			FullName:    fake.FullName(),
			PhoneNumber: fakePhoneNumber(),
		}

		if err := r.Create(context.TODO(), m); err != nil {
			return err
		}
	}

	return nil
}

func addCustomer(db *gorm.DB) error {
	r := repositories.NewCustomerRepository(db)
	for i := 0; i < 20; i++ {
		m := &models.Customer{
			FullName:    fake.FullName(),
			PhoneNumber: fakePhoneNumber(),
		}

		if err := r.Create(context.TODO(), m); err != nil {
			return err
		}
	}

	return nil
}

func addVendor(db *gorm.DB) error {
	r := repositories.NewVendorRepository(db)
	for i := 0; i < 20; i++ {
		m := &models.Vendor{
			Name:        fake.Brand(),
			Address:     fake.StreetAddress(),
			PhoneNumber: fakePhoneNumber(),
		}

		if err := r.Create(context.TODO(), m); err != nil {
			return err
		}
	}

	return nil
}

func addOrder(db *gorm.DB) error {
	r := repositories.NewOrderRepository(db)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		m := &models.Order{
			CustomerID:   rand.Intn(15) + 1,
			VendorID:     rand.Intn(15) + 1,
			DeliveryTime: time.Now().Add(time.Duration(rand.Intn(20)+1) * time.Minute),
		}

		if err := r.Create(context.TODO(), m); err != nil {
			return err
		}
	}

	return nil
}

func addTrip(db *gorm.DB) error {
	r := repositories.NewTripRepository(db)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		m := &models.Trip{
			CourierID: rand.Intn(15) + 1,
			OrderID:   rand.Intn(10) + 1,
			Status:    models.TripStatus(rand.Intn(3)),
		}

		if err := r.Create(context.TODO(), m); err != nil {
			return err
		}
	}

	return nil
}

func addDelayReport(db *gorm.DB) error {
	r := repositories.NewDelayReportRepository(db)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		m := &models.DelayReport{
			OrderID:    rand.Intn(10) + 1,
			AgentID:    rand.Intn(10) + 1,
			ReportTime: time.Now(),
			Status:     models.ReportStatus(rand.Intn(2)),
		}

		if err := r.Create(context.TODO(), m); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	log.Println("Connecting to database...")
	db := Connect2DB()
	log.Println("connected to database successfully!")

	log.Println("Start seeding...")
	if err := addAgents(db); err != nil {
		log.Fatalf("error on creating agents: %v", err)
	}

	if err := addCourier(db); err != nil {
		log.Fatalf("error on creating couriers: %v", err)
	}

	if err := addCustomer(db); err != nil {
		log.Fatalf("error on creating customers: %v", err)
	}

	if err := addVendor(db); err != nil {
		log.Fatalf("error on creating vendors: %v", err)
	}

	if err := addOrder(db); err != nil {
		log.Fatalf("error on creating orders: %v", err)
	}

	if err := addTrip(db); err != nil {
		log.Fatalf("error on creating trips: %v", err)
	}

	if err := addDelayReport(db); err != nil {
		log.Fatalf("error on creating delay reports: %v", err)
	}

	log.Println("seeding completed successfully!")
}

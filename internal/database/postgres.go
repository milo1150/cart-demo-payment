package database

import (
	"fmt"
	"log"
	"os"

	"github.com/milo1150/cart-demo-payment/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// Load env
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	pwd := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_DOCKER_PORT")
	timezone := os.Getenv("TIMEZONE")
	conn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=%v", host, user, pwd, name, port, timezone)

	// Connect db
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[error] failed to connect to database: %v\n", err)
	}

	return db
}

func RunAutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.PaymentOrder{})
}

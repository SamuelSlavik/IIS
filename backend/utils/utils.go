package utils

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env variables!")
	}
}

var DB *gorm.DB

func Conn() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
}

type CustomDate struct {
	time.Time
}

func (d *CustomDate) UnmarshalJSON(b []byte) error {
	// Define the date format you expect in your JSON data
	dateFormat := "\"2006-01-02\""

	// Parse the JSON data using the specified format
	dateStr := string(b)
	parsedDate, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return err
	}

	d.Time = parsedDate
	return nil
}

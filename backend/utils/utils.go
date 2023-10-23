package utils

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conn() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Prague"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
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

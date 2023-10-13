package utils

import (
	"github.com/AdamPekny/IIS/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conn() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func Migrate_all() {
	db, _ := Conn()

	// Migrate User models
	db.AutoMigrate(&models.User{}, &models.UserType{})

	// Migrate Vehicle models
	db.AutoMigrate(&models.Vehicle{}, &models.VehicleType{}, &models.Line{}, &models.Stop{})

	// Migrate Line and Stop models
	db.AutoMigrate(&models.Line{}, &models.Stop{}) // $models.TimeBetween{}

	// Migrate Connection models
	db.AutoMigrate(&models.Connection{})

}

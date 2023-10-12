package api

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "github.com/AdamPekny/IIS/backend/models"
)

func Router() {
	db, _ := conn()
	fmt.Print(db.Table("information_schema.tables"))
	db.AutoMigrate(&models.User{}, &models.UserType{}, &models.Line{}, &models.VehicleType{}, &models.Vehicle{})
}

func conn() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

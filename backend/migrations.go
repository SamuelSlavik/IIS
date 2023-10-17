package api

import (
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
)

func Migrate_all() {
	db, _ := utils.Conn()

	// Migrate User models
	db.AutoMigrate(&models.User{}, &models.UserType{})

	// Migrate Vehicle models
	db.AutoMigrate(&models.Line{}, &models.Stop{}, &models.Connection{},
		&models.Vehicle{}, &models.VehicleType{}, models.Segment{})

}

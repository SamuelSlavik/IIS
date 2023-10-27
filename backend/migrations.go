package api

import (
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
)

func Migrate_all() {
	// Migrate User models
	utils.DB.AutoMigrate(&models.User{})

	// Migrate Vehicle models
	utils.DB.AutoMigrate(&models.Line{}, &models.Stop{}, &models.Connection{},
		&models.Vehicle{}, &models.VehicleType{}, models.Segment{})
}

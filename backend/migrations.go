package api

import (
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
)

func Migrate_all() {
	// Migrate User models
	utils.DB.AutoMigrate(&models.User{})

	// Migrate Vehicle models
	utils.DB.AutoMigrate(&models.VehicleType{}, &models.Vehicle{})

	utils.DB.AutoMigrate(&models.Stop{}, &models.Line{}, &models.Segment{}, &models.Connection{})
}

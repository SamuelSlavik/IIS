package api

import (
	"github.com/AdamPekny/IIS/backend/userauth"
	"github.com/AdamPekny/IIS/backend/utils"
)

func Migrate_all() {
	db, _ := utils.Conn()

	// Migrate User models
	db.AutoMigrate(&userauth.User{}, &userauth.UserType{})

	// Migrate Vehicle models
	// db.AutoMigrate(&models.Vehicle{}, &models.VehicleType{}, &models.Line{}, &models.Stop{})

	// Migrate Line and Stop models
	// db.AutoMigrate(&models.Line{}, &models.Stop{}) // $models.TimeBetween{}

	// Migrate Connection models
	// db.AutoMigrate(&models.Connection{})

}

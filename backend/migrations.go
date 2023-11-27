package api

import (
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
	"golang.org/x/crypto/bcrypt"
)

func Migrate_all() {
	// Migrate User models
	utils.DB.AutoMigrate(&models.User{})

	// Migrate Vehicle models
	utils.DB.AutoMigrate(&models.VehicleType{}, &models.Vehicle{})

	// Migrate Maintenance models
	utils.DB.AutoMigrate(&models.MalfunctionReport{}, &models.MaintenanceRequest{}, &models.MaintenanceReport{})

	utils.DB.AutoMigrate(&models.Stop{}, &models.Line{}, &models.Segment{}, &models.Connection{})

	var users []models.User
	result := utils.DB.Where("role = ?", string(models.AdminRole)).Find(&users)
	
	if result.Error != nil {
		return
	}

	pwd_hash, err := bcrypt.GenerateFromPassword([]byte("DmiInbN5"), 14)
	if err != nil {
		return
	}

	if result.RowsAffected == 0 {
		utils.DB.Create(&models.User{
			FirstName: "user",
			LastName: "admin",
			Email: "user@admin.com",
			BirthDate: time.Now(),
			Password: string(pwd_hash),
			Role: models.AdminRole,
		})
	}
}

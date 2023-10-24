package serializers

import (
	models "github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/validators"
)

type VehicleSerializer struct {
	//Registration string `binding:"required"`
	Capacity      uint   `binding:"required"`
	Registration  string `binding:"required"`
	Brand         string
	ImageData     []byte
	Type          string `binding:"required"`
	LineName      string
	ValidatorErrs []validators.ValidatorErr
}

func (vehicle VehicleSerializer) Create_model() *models.Vehicle {
	vehicle_model := &models.Vehicle{}

	//copy data
	vehicle_model.Capacity = vehicle.Capacity
	vehicle_model.Registration = vehicle.Registration
	vehicle_model.Brand = vehicle.Brand
	vehicle_model.ImageData = vehicle.ImageData
	vehicle_model.VehicleTypeName = vehicle.Type
	if vehicle.LineName == "" {
		vehicle_model.LineName = nil
	} else {
		vehicle_model.LineName = &vehicle.LineName
	}

	return vehicle_model
}

func (vehicle *VehicleSerializer) Valid() bool {
	validators.Registration_validator(vehicle.Registration, &vehicle.ValidatorErrs)
	return len(vehicle.ValidatorErrs) == 0
}

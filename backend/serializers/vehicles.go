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
	LineName      *string
	ValidatorErrs []validators.ValidatorErr
}

func (vehicle VehicleSerializer) Create_model() (vehicle_model *models.Vehicle) {
	vehicle_model = &models.Vehicle{
		Capacity:        vehicle.Capacity,
		Registration:    vehicle.Registration,
		Brand:           vehicle.Brand,
		ImageData:       vehicle.ImageData,
		VehicleTypeName: vehicle.Type,
		LineName:        vehicle.LineName,
	}
	return
}

func (vehicle *VehicleSerializer) Valid() bool {
	validators.Registration_validator(vehicle.Registration, &vehicle.ValidatorErrs)
	return len(vehicle.ValidatorErrs) == 0
}

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
	ValidatorErrs []validators.ValidatorErr
}

type VehicleGetSerializer struct {
	Registration string
	Capacity     uint
	Brand        string
	//image lol
	Type            string
	LastMaintenance LastMaintenance
}
type LastMaintenance struct {
	Status string
	Date   string
}

func (vehicle VehicleSerializer) Create_model() (vehicle_model *models.Vehicle) {
	vehicle_model = &models.Vehicle{
		Capacity:        vehicle.Capacity,
		Registration:    vehicle.Registration,
		Brand:           vehicle.Brand,
		ImageData:       vehicle.ImageData,
		VehicleTypeName: vehicle.Type,
	}
	return
}

func (vehicle *VehicleSerializer) Valid() bool {
	validators.Registration_validator(vehicle.Registration, &vehicle.ValidatorErrs)
	return len(vehicle.ValidatorErrs) == 0
}

type VehicleMaintenanceSerializer struct {
	Registration    string
	Brand           string
	VehicleTypeName string
	ValidatorErrs   []validators.ValidatorErr
}

func (v *VehicleMaintenanceSerializer) Valid() bool {
	return true
}

func (v *VehicleMaintenanceSerializer) FromModel(vehicle_model *models.Vehicle) (err error) {
	v.Registration = vehicle_model.Registration
	v.Brand = vehicle_model.Brand
	v.VehicleTypeName = vehicle_model.VehicleTypeName

	return nil
}

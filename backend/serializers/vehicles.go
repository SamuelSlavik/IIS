package serializers

import (
	models "github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/validators"
)

type VehicleSerializer struct {
	//Registration string `binding:"required"`
	Capacity     uint   `binding:"required"`
	Registration string `binding:"required"`
	Brand        string
	//ImageData     []byte
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
type VehicleTypeSerializer struct {
	Type   string
	Active bool
	//IconPath string `binding:"required"`
}
type VehicleTypeCreateSerializer struct {
	Type string `binding:"required"`
	//IconPath string `binding:"required"`
}

type LastMaintenance struct {
	Status string
	Date   string
}

type VehicleUpdateSerializer struct {
	Capacity uint `binding:"required"`
	Brand    string
	//image lol
	Type          string `binding:"required"`
	ValidatorErrs []validators.ValidatorErr
}

func (vehicle VehicleSerializer) Create_model() (vehicle_model *models.Vehicle) {
	vehicle_model = &models.Vehicle{
		Capacity:     vehicle.Capacity,
		Registration: vehicle.Registration,
		Brand:        vehicle.Brand,
		//ImageData:       vehicle.ImageData,
		VehicleTypeName: vehicle.Type,
	}
	return
}

func (vehicle *VehicleUpdateSerializer) Valid() bool {
	validators.Vehicle_type_validator(vehicle.Type, &vehicle.ValidatorErrs)
	return len(vehicle.ValidatorErrs) == 0
}

func (vehicle *VehicleSerializer) Valid() bool {
	validators.Registration_validator(vehicle.Registration, &vehicle.ValidatorErrs)
	validators.Vehicle_type_validator(vehicle.Type, &vehicle.ValidatorErrs)
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

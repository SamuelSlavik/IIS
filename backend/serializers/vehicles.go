// package serializers holds structures and functions for serializing data
// this file contains serializers for vehicles
package serializers

import (
	models "github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/validators"
)

// VehicleSerializer is used to serialize data about vehicle
// it is used in POST request to create a new vehicle
type VehicleSerializer struct {
	//Registration string `binding:"required"`
	Capacity     uint   `binding:"required"`
	Registration string `binding:"required"`
	Brand        string
	//ImageData     []byte
	Type          string `binding:"required"`
	ValidatorErrs []validators.ValidatorErr
}

// VehicleGetSerializer is used to serialize data from database about vehicle
// it is used in GET request to get data about vehicle
type VehicleGetSerializer struct {
	Registration string
	Capacity     uint
	Brand        string
	//image lol
	Type            string
	LastMaintenance LastMaintenance
}

// VehicleTypeSerializer is used to serialize data about vehicle type
// it is used in GET request to get data about vehicle type
type VehicleTypeSerializer struct {
	Type   string
	Active bool
	//IconPath string `binding:"required"`
}

// VehicleTypeGetSerializer is used to serialize data about vehicle
// it is used in POST request to create a new vehicle
type VehicleTypeCreateSerializer struct {
	Type string `binding:"required"`
	//IconPath string `binding:"required"`
}

// LastMaintenance is used to serialize data about last maintenance
type LastMaintenance struct {
	Status string
	Date   string
}

// VehicleUpdateSerializer is used to serialize data about vehicle
// it is used in PUT request to update a vehicle
type VehicleUpdateSerializer struct {
	Capacity uint `binding:"required"`
	Brand    string
	//image lol
	Type          string `binding:"required"`
	ValidatorErrs []validators.ValidatorErr
}

// Create_model creates a new vehicle model from serializer
// it is used in POST request to create a new vehicle model
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

// Valid validates the data from serializer
func (vehicle *VehicleUpdateSerializer) Valid() bool {
	validators.Vehicle_type_validator(vehicle.Type, &vehicle.ValidatorErrs)
	return len(vehicle.ValidatorErrs) == 0
}

// Valid validates the data from serializer
func (vehicle *VehicleSerializer) Valid() bool {
	validators.Registration_validator(vehicle.Registration, &vehicle.ValidatorErrs)
	validators.Vehicle_type_validator(vehicle.Type, &vehicle.ValidatorErrs)
	return len(vehicle.ValidatorErrs) == 0
}

// VehicleMaintenanceSerializer is used to serialize data about vehicle
// it is used in GET request to get data about vehicle
type VehicleMaintenanceSerializer struct {
	Registration    string
	Brand           string
	VehicleTypeName string
	ValidatorErrs   []validators.ValidatorErr
}

func (v *VehicleMaintenanceSerializer) Valid() bool {
	return true
}

// FromModel loads data from model into serializer
func (v *VehicleMaintenanceSerializer) FromModel(vehicle_model *models.Vehicle) (err error) {
	v.Registration = vehicle_model.Registration
	v.Brand = vehicle_model.Brand
	v.VehicleTypeName = vehicle_model.VehicleTypeName

	return nil
}

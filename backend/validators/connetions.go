package validators

import (
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
)

func Line_name_validator(name string, validator_errs *[]ValidatorErr) {
	res := utils.DB.Where("name = ?", name).First(&models.Line{})
	if res.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"DatabaseErr", res.Error.Error()})
		return
	}
	if res.RowsAffected == 0 {
		*validator_errs = append(*validator_errs, ValidatorErr{"CreatingConnErr", "Line with given name does not exist"})
	}
}

func Vehicle_registration_validator(registration string, validator_errs *[]ValidatorErr) {
	res := utils.DB.Where("registration = ?", registration).First(&models.Vehicle{})
	if res.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"DatabaseErr", res.Error.Error()})
		return
	}
	if res.RowsAffected == 0 {
		*validator_errs = append(*validator_errs, ValidatorErr{"CreatingConnErr", "Vehicle with given registration does not exist"})
	}
}

func Driver_id_validator(id uint, validator_errs *[]ValidatorErr) {
	user := models.User{}
	res := utils.DB.Where("id = ?", id).First(&user)
	if res.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"DatabaseErr", res.Error.Error()})
		return
	}
	if res.RowsAffected == 0 {
		*validator_errs = append(*validator_errs, ValidatorErr{"CreatingConnErr", "Driver with given id does not exist"})
		return
	}
	if user.Role != "driver" {
		*validator_errs = append(*validator_errs, ValidatorErr{"CreatingConnErr", "User with given id is not a driver"})
		return
	}
}

func Vehicle_availability(registration string, departure_time string, validator_errs *[]ValidatorErr) {
	var vehicle models.Vehicle
	timeString := departure_time
	timeObject, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"TimeParse", err.Error()})
		return
	}
	res := utils.DB.Preload("Connections").Where("registration = ?", registration).First(&vehicle)
	if res.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"DatabaseErr", res.Error.Error()})
		return
	}
	for _, connection := range vehicle.Connections {
		if connection.DepartureTime.Before(timeObject) || connection.ArrivalTime.After(timeObject) {
			*validator_errs = append(*validator_errs, ValidatorErr{"VehicleAvailability", "Vehicle is not available at given time"})
			return
		}
	}

}

// todo optimalizovat 2x pristup do db
func Driver_availability(driverID uint, departure_time string, validator_errs *[]ValidatorErr) {
	var user models.User
	timeString := departure_time
	timeObject, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"TimeParse", err.Error()})
		return
	}
	res := utils.DB.Preload("Connetions").Where("id = ?", driverID).Find(&user)
	if res.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"DatabaseErr", res.Error.Error()})
		return
	}
	for _, connection := range user.Connections {
		if connection.DepartureTime.Before(timeObject) || connection.ArrivalTime.After(timeObject) {
			*validator_errs = append(*validator_errs, ValidatorErr{"DriverAvailability", "Driver is not available at given time"})
			return
		}
	}

}

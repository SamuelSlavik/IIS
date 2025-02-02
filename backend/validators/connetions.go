// package validators contains functions for validating recieved data
// this file contains validators for connections
package validators

import (
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
)

// Line_name_validator validates line name
// loads any errors into validator_errs
func Line_name_validator(name string, validator_errs *[]ValidatorErr) {
	res := utils.DB.Where("name = ?", name).Find(&models.Line{})
	if res.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"DatabaseErr", res.Error.Error()})
		return
	}
	if res.RowsAffected == 0 {
		*validator_errs = append(*validator_errs, ValidatorErr{"CreatingConnErr", "Line with given name does not exist"})
	}
}

// Vehicle_registration_validator validates if vehicle registration exists
func Vehicle_registration_validator(registration *string, validator_errs *[]ValidatorErr) {
	if registration == nil {
		return
	}

	res := utils.DB.Where("registration = ?", registration).Find(&models.Vehicle{})
	if res.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"DatabaseErr", res.Error.Error()})
		return
	}
	if res.RowsAffected == 0 {
		*validator_errs = append(*validator_errs, ValidatorErr{"CreatingConnErr", "Vehicle with given registration does not exist"})
	}
}

// Driver_id_validator validates if driver exists and if it is a driver
func Driver_id_validator(id *uint, validator_errs *[]ValidatorErr) {
	if id == nil {
		return
	}
	user := models.User{}
	res := utils.DB.Where("id = ?", id).Find(&user)
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

// Vehicle_availability checks if vehicle is available at given time
func Vehicle_availability(id int, registration *string, departure_time string, arrival_time time.Time, NumberOfDays int, validator_errs *[]ValidatorErr) {
	if registration == nil {
		return
	}
	if NumberOfDays < 1 {
		*validator_errs = append(*validator_errs, ValidatorErr{"CreatingConnErr", "Number of days must be greater than 0"})
		return
	}
	var vehicle models.Vehicle
	timeString := departure_time
	timeObject, err := time.Parse("2006-01-02 15:04", timeString)
	if err != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"TimeParse", err.Error()})
		return
	}
	res := utils.DB.Preload("Connections").Where("registration = ?", registration).Find(&vehicle)
	if res.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"DatabaseErr", res.Error.Error()})
		return
	}
	for i := 0; i < int(NumberOfDays); i++ {
		timeObject = timeObject.AddDate(0, 0, i)
		arrival_time = arrival_time.AddDate(0, 0, i)
		for _, connection := range vehicle.Connections {
			if id != -1 && int(connection.ID) == id {
				continue
			}
			if (connection.DepartureTime.Before(timeObject) || connection.DepartureTime.Equal(timeObject)) && (connection.ArrivalTime.After(timeObject) || connection.ArrivalTime.Equal(timeObject)) {
				*validator_errs = append(*validator_errs, ValidatorErr{"VehicleAvailability", "Vehicle is not available at given time"})
				return
			}
			if (connection.DepartureTime.Before(arrival_time) || connection.DepartureTime.Equal(arrival_time)) && (connection.ArrivalTime.After(arrival_time) || connection.ArrivalTime.Equal(arrival_time)) {
				*validator_errs = append(*validator_errs, ValidatorErr{"VehicleAvailability", "Vehicle is not available at given time"})
				return
			}
		}
	}

}

// Driver_availability checks if driver is available at given time
func Driver_availability(id int, driverID *uint, departure_time string, arrival_time time.Time, NumberOfDays int, validator_errs *[]ValidatorErr) {
	if driverID == nil {
		return
	}
	if NumberOfDays < 1 {
		*validator_errs = append(*validator_errs, ValidatorErr{"CreatingConnErr", "Number of days must be greater than 0"})
		return
	}
	var user models.User
	timeString := departure_time
	timeObject, err := time.Parse("2006-01-02 15:04", timeString)
	if err != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"TimeParse", err.Error()})
		return
	}
	res := utils.DB.Preload("Connections").Where("id = ?", driverID).Find(&user)
	if res.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{"DatabaseErr", res.Error.Error()})
		return
	}
	for i := 0; i < int(NumberOfDays); i++ {
		timeObject = timeObject.AddDate(0, 0, i)
		arrival_time = arrival_time.AddDate(0, 0, i)
		for _, connection := range user.Connections {
			if id != -1 && int(connection.ID) == id {
				continue
			}
			if (connection.DepartureTime.Before(timeObject) || connection.DepartureTime.Equal(timeObject)) && (connection.ArrivalTime.After(timeObject) || connection.ArrivalTime.Equal(timeObject)) {
				*validator_errs = append(*validator_errs, ValidatorErr{"DriverAvailability", "Driver is not available at given time"})
				return
			}
			if (connection.DepartureTime.Before(arrival_time) || connection.DepartureTime.Equal(arrival_time)) && (connection.ArrivalTime.After(arrival_time) || connection.ArrivalTime.Equal(arrival_time)) {
				*validator_errs = append(*validator_errs, ValidatorErr{"DriverAvailability", "Driver is not available at given time"})
				return
			}
		}
	}

}

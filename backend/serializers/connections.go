package serializers

import (
	"fmt"
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/AdamPekny/IIS/backend/validators"
)

type ConnectionSerializer struct {
	ConnectionID  uint
	LineName      string
	DepartureTime string
	ArrivalTime   string
	Direction     bool
	InitialStop   string
	FinalStop     string
	VehicleReg    *string
	DriverID      *uint
}

type ConnectionCreateSerializer struct {
	LineName      string `binding:"required"`
	DepartureTime string `binding:"required"`
	VehicleReg    *string
	Direction     bool
	DriverID      *uint
	NumberOfDays  int       `binding:"required"`
	ArrivalTime   time.Time //neplnit z fe
	ValidatorErrs []validators.ValidatorErr
}

type ConnectionAssignSerializer struct {
	DriverID      *uint
	VehicleReg    *string
	NumberOfDays  int       `binding:"required"`
	DepartureTime string    //neplnit z fe
	ArrivalTime   time.Time //neplnit z fe
	ValidatorErrs []validators.ValidatorErr
}

type ConnectionUpdateSerializer struct {
	LineName      string
	DepartureTime string
	Direction     bool
	ArrivalTime   time.Time //neplnit z fe
	DriverID      *uint
	VehicleReg    *string
	NumberOfDays  int `binding:"required"`
	ValidatorErrs []validators.ValidatorErr
}

type ConnectionDeleteSerializer struct {
	NumberOfDays int `binding:"required"`
}

type ConnectionDetailsSerializer struct {
	ID        uint
	LineName  string
	Type      string
	ListStops *[]StopInConnection
}

type StopInConnection struct {
	DepartureTime string
	StopName      string
}

func Get_arrival_time(dep_time time.Time, line_name string) (arrival_time time.Time) {
	line := models.Line{}
	utils.DB.Preload("Segments").First(&line, "name=?", line_name)
	var duration time.Duration
	for _, segment := range line.Segments {
		duration += time.Minute * time.Duration(segment.Time)
	}
	arrival_time = dep_time.Add(duration)
	return
}

func (conn *ConnectionCreateSerializer) Valid() bool {
	validators.Line_name_validator(conn.LineName, &conn.ValidatorErrs)
	validators.Vehicle_registration_validator(conn.VehicleReg, &conn.ValidatorErrs)
	validators.Driver_id_validator(conn.DriverID, &conn.ValidatorErrs)
	if len(conn.ValidatorErrs) != 0 {
		return false
	}

	dep_time, _ := time.Parse("2006-01-02 15:04", conn.DepartureTime)
	conn.ArrivalTime = Get_arrival_time(dep_time, conn.LineName)
	validators.Vehicle_availability(-1, conn.VehicleReg, conn.DepartureTime, conn.ArrivalTime, conn.NumberOfDays, &conn.ValidatorErrs)
	validators.Driver_availability(-1, conn.DriverID, conn.DepartureTime, conn.ArrivalTime, conn.NumberOfDays, &conn.ValidatorErrs)
	return len(conn.ValidatorErrs) == 0

}
func (conn *ConnectionAssignSerializer) Valid(id int) bool {
	validators.Vehicle_registration_validator(conn.VehicleReg, &conn.ValidatorErrs)
	validators.Driver_id_validator(conn.DriverID, &conn.ValidatorErrs)
	if len(conn.ValidatorErrs) != 0 {
		return false
	}
	validators.Vehicle_availability(id, conn.VehicleReg, conn.DepartureTime, conn.ArrivalTime, 1, &conn.ValidatorErrs)
	validators.Driver_availability(id, conn.DriverID, conn.DepartureTime, conn.ArrivalTime, 1, &conn.ValidatorErrs)
	return len(conn.ValidatorErrs) == 0
}

func (conn ConnectionCreateSerializer) CreateModel() (connection_model []models.Connection, err error) {
	dep_time, err := time.Parse("2006-01-02 15:04", conn.DepartureTime)
	if err != nil {
		return
	}
	for i := 0; i < int(conn.NumberOfDays); i++ {
		res := utils.DB.Where("line_name = ? AND departure_time = ?", conn.LineName, dep_time).Find(&models.Connection{})
		if res.Error != nil {
			err = res.Error
			return
		}
		if res.RowsAffected != 0 {
			err = fmt.Errorf("Connection at same time for this line already exists")
			return
		}
		connection_model = append(connection_model, models.Connection{
			LineName:            conn.LineName,
			DepartureTime:       dep_time,
			ArrivalTime:         conn.ArrivalTime,
			VehicleRegistration: conn.VehicleReg,
			Direction:           conn.Direction,
			DriverID:            conn.DriverID,
		})
		dep_time = dep_time.Add(time.Hour * 24)
		conn.ArrivalTime = conn.ArrivalTime.Add(time.Hour * 24)
	}
	return
}

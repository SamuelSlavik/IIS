package serializers

import (
	"fmt"
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/AdamPekny/IIS/backend/validators"
)

type ConnectionSerializer struct {
	ID            uint
	LineName      string
	DepartureTime string
	ArrivalTime   string
	VehicleReg    *string
	DriverID      *uint
	Dirrection    bool
}

type ConnectionLineSerializer struct {
	ConnectionID  uint
	LineName      string
	DepartureTime string
	ArrivalTime   string
	Dirrection    bool
	InitialStop   string
	FinalStop     string
}

type ConnectionCreateSerializer struct {
	LineName      string `binding:"required"`
	DepartureTime string `binding:"required"`
	VehicleReg    *string
	Dirrection    bool `binding:"required"`
	DriverID      *uint
	NumberOfDays  int       `binding:"required"`
	ArrivalTime   time.Time //neplnit z fe
	ValidatorErrs []validators.ValidatorErr
}

type ConnectionAssignSerializer struct {
	DriverID      *uint
	VehicleReg    *string
	DepartureTime string    //neplnit z fe
	ArrivalTime   time.Time //neplnit z fe
	ValidatorErrs []validators.ValidatorErr
}

type ConnectionUpdateSerializer struct {
	LineName      string
	DepartureTime string
	Dirrection    bool
	ArrivalTime   time.Time //neplnit z fe
	DriverID      *uint
	VehicleReg    *string
	ValidatorErrs []validators.ValidatorErr
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
		fmt.Print(segment.Time)
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

	dep_time, _ := time.Parse("2006-01-02 15:04:05", conn.DepartureTime)
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

func (conn ConnectionCreateSerializer) CreateModel() (connection_model []models.Connection) {
	dep_time, _ := time.Parse("2006-01-02 15:04:05", conn.DepartureTime) //todo lolik error
	for i := 0; i < int(conn.NumberOfDays); i++ {
		connection_model = append(connection_model, models.Connection{
			LineName:            conn.LineName,
			DepartureTime:       dep_time,
			ArrivalTime:         conn.ArrivalTime,
			VehicleRegistration: conn.VehicleReg,
			Dirrection:          conn.Dirrection,
			DriverID:            conn.DriverID,
		})
		dep_time = dep_time.Add(time.Hour * 24)
		conn.ArrivalTime = conn.ArrivalTime.Add(time.Hour * 24)
	}
	return
}

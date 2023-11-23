package serializers

import (
	"fmt"
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/AdamPekny/IIS/backend/validators"
)

type ConnectionSerializer struct {
	ID        uint
	LineName  string
	Type      string
	ListStops *[]StopInConnection
}

type StopInConnection struct {
	DepartureTime string
	StopName      string
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
	ValidatorErrs []validators.ValidatorErr
	NumberOfDays  int `binding:"required"`
}

func (conn *ConnectionCreateSerializer) Valid() bool {
	validators.Line_name_validator(conn.LineName, &conn.ValidatorErrs)
	validators.Vehicle_registration_validator(conn.VehicleReg, &conn.ValidatorErrs)
	validators.Driver_id_validator(conn.DriverID, &conn.ValidatorErrs)
	if len(conn.ValidatorErrs) != 0 {
		return false
	}
	validators.Vehicle_availability(conn.VehicleReg, conn.DepartureTime, conn.NumberOfDays, &conn.ValidatorErrs)
	validators.Driver_availability(conn.DriverID, conn.DepartureTime, conn.NumberOfDays, &conn.ValidatorErrs)
	return len(conn.ValidatorErrs) == 0

}
func (conn ConnectionCreateSerializer) CreateModel() (connection_model []models.Connection) {
	line := models.Line{}
	utils.DB.Preload("Segments").First(&line, "name=?", conn.LineName)
	var duration time.Duration
	for _, segment := range line.Segments {
		fmt.Print(segment.Time)
		duration += time.Minute * time.Duration(segment.Time)
	}
	dep_time, _ := time.Parse("2006-01-02 15:04:05", conn.DepartureTime) //todo lolik error
	for i := 0; i < int(conn.NumberOfDays); i++ {
		connection_model = append(connection_model, models.Connection{
			LineName:            conn.LineName,
			DepartureTime:       dep_time,
			ArrivalTime:         dep_time.Add(duration),
			VehicleRegistration: conn.VehicleReg,
			Dirrection:          conn.Dirrection,
			DriverID:            conn.DriverID,
		})
		dep_time = dep_time.Add(time.Hour * 24)
	}
	return
}

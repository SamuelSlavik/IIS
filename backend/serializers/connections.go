package serializers

import "github.com/AdamPekny/IIS/backend/validators"

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
	VehicleReg    string
	Dirrection    bool `binding:"required"`
	DriverID      uint
	ValidatorErrs []validators.ValidatorErr
}

func (conn *ConnectionCreateSerializer) Valid() bool {
	validators.Line_name_validator(conn.LineName, &conn.ValidatorErrs)
	validators.Vehicle_registration_validator(conn.VehicleReg, &conn.ValidatorErrs)
	validators.Driver_id_validator(conn.DriverID, &conn.ValidatorErrs)
	if len(conn.ValidatorErrs) != 0 {
		return false
	}
	validators.Vehicle_availability(conn.VehicleReg, conn.DepartureTime, &conn.ValidatorErrs)
	validators.Driver_availability(conn.DriverID, conn.DepartureTime, &conn.ValidatorErrs)
	return len(conn.ValidatorErrs) == 0

}

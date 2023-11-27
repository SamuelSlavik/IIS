package views

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/serializers"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/AdamPekny/IIS/backend/validators"
	"github.com/gin-gonic/gin"
)

func ListConnections(ctx *gin.Context) {
	var connections []serializers.ConnectionSerializer
	var connection_models []models.Connection
	var err error
	err = utils.DB.Find(&connection_models).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	for _, model := range connection_models {
		line := models.Line{}
		err = utils.DB.First(&line, "name=?", model.LineName).Error
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		connection := serializers.ConnectionSerializer{
			ConnectionID:  model.ID,
			LineName:      model.LineName,
			DepartureTime: model.DepartureTime.Format("2006-01-02 15:04"),
			ArrivalTime:   model.ArrivalTime.Format("2006-01-02 15:04"),
			InitialStop:   line.InitialStop,
			FinalStop:     line.FinalStop,
			DriverID:      model.DriverID,
			VehicleReg:    model.VehicleRegistration,
			Direction:     model.Direction,
		}
		if connection.Direction == true {
			connection.InitialStop = line.FinalStop
			connection.FinalStop = line.InitialStop
		}
		connections = append(connections, connection)
	}
	ctx.IndentedJSON(http.StatusOK, connections)
}

// for unregistered
func GetDetailOfConnection(ctx *gin.Context) {
	id := ctx.Param("id")
	var connection_model models.Connection
	var err error
	err = utils.DB.Where("driver_id IS NOT NULL AND vehicle_registration IS NOT NULL").First(&connection_model, "id=?", id).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	connection := serializers.ConnectionDetailsSerializer{
		ID:       connection_model.ID,
		LineName: connection_model.LineName,
	}
	connection.ListStops, err = getStops(connection.ID)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	var vehicle models.Vehicle
	err = utils.DB.First(&vehicle, "id=?", connection_model.VehicleRegistration).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	connection.Type = vehicle.VehicleTypeName
	ctx.IndentedJSON(http.StatusOK, connection)
}

func getStops(lineID uint) (*[]serializers.StopInConnection, error) {
	var connection models.Connection
	stops := []serializers.StopInConnection{}
	if err := utils.DB.First(&connection, lineID).Error; err != nil {
		return nil, err
	}
	var line models.Line
	if err := utils.DB.Model(&line).Preload("Segments").First(&line, "Name = ?", connection.LineName).Error; err != nil {
		return nil, err
	}
	dep_time := connection.DepartureTime
	for i := 0; i < len(line.Segments); i++ {
		if line.Segments[i].StopName2 == line.FinalStop {
			stops = append(stops, serializers.StopInConnection{
				StopName:      line.Segments[i].StopName1,
				DepartureTime: dep_time.Format("15:04"),
			})
			dep_time = dep_time.Add(time.Minute * time.Duration(line.Segments[i].Time))
			stops = append(stops, serializers.StopInConnection{
				StopName:      line.Segments[i].StopName2,
				DepartureTime: dep_time.Format("15:04"),
			})
			break
		}
		stops = append(stops, serializers.StopInConnection{
			StopName:      line.Segments[i].StopName1,
			DepartureTime: dep_time.Format("15:04"),
		})
		dep_time = dep_time.Add(time.Minute * time.Duration(line.Segments[i].Time))

	}
	return &stops, nil
}

func GetConnectionById(ctx *gin.Context) {
	id := ctx.Param("id")
	var connection_model models.Connection
	var err error
	err = utils.DB.First(&connection_model, "id=?", id).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	line := models.Line{}
	err = utils.DB.First(&line, "name=?", connection_model.LineName).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	connection := serializers.ConnectionSerializer{
		ConnectionID:  connection_model.ID,
		LineName:      connection_model.LineName,
		DepartureTime: connection_model.DepartureTime.Format("2006-01-02 15:04"),
		ArrivalTime:   connection_model.ArrivalTime.Format("2006-01-02 15:04"),
		InitialStop:   line.InitialStop,
		FinalStop:     line.FinalStop,
		DriverID:      connection_model.DriverID,
		VehicleReg:    connection_model.VehicleRegistration,
		Direction:     connection_model.Direction,
	}
	if connection.Direction == true {
		connection.InitialStop = line.FinalStop
		connection.FinalStop = line.InitialStop
	}
	ctx.IndentedJSON(http.StatusOK, connection)
}

func ListConnectionsByLine(ctx *gin.Context) {
	line := ctx.Param("line")
	var connection_models []models.Connection
	var err error
	line_model := models.Line{}
	err = utils.DB.First(&line_model, "name=?", line).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	err = utils.DB.Find(&connection_models, "line_name=?", line).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	connections := []serializers.ConnectionSerializer{}
	for _, model := range connection_models {
		connection := serializers.ConnectionSerializer{
			ConnectionID:  model.ID,
			LineName:      model.LineName,
			ArrivalTime:   model.ArrivalTime.Format("2006-01-02 15:04"),
			DepartureTime: model.DepartureTime.Format("2006-01-02 15:04"),
			Direction:     model.Direction,
			InitialStop:   line_model.InitialStop,
			FinalStop:     line_model.FinalStop,
			VehicleReg:    model.VehicleRegistration,
			DriverID:      model.DriverID,
		}
		if connection.Direction == true {
			connection.InitialStop = line_model.FinalStop
			connection.FinalStop = line_model.InitialStop
		}
		connections = append(connections, connection)
	}
	ctx.IndentedJSON(http.StatusOK, connections)
}

func ListConnectionsByLineAndDate(ctx *gin.Context) {
	line := ctx.Param("line")
	date := ctx.Param("date")
	var connection_models []models.Connection
	var err error
	line_model := models.Line{}
	err = utils.DB.First(&line_model, "name=?", line).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	err = utils.DB.Find(&connection_models, "line_name=? AND departure_time BETWEEN ? AND ? ", line, date, date+" 23:59:59").Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	connections := []serializers.ConnectionSerializer{}
	for _, model := range connection_models {
		connection := serializers.ConnectionSerializer{
			ConnectionID:  model.ID,
			LineName:      model.LineName,
			ArrivalTime:   model.ArrivalTime.Format("2006-01-02 15:04"),
			DepartureTime: model.DepartureTime.Format("2006-01-02 15:04"),
			Direction:     model.Direction,
			InitialStop:   line_model.InitialStop,
			FinalStop:     line_model.FinalStop,
			VehicleReg:    model.VehicleRegistration,
			DriverID:      model.DriverID,
		}
		if connection.Direction == true {
			connection.InitialStop = line_model.FinalStop
			connection.FinalStop = line_model.InitialStop
		}
		connections = append(connections, connection)
	}
	ctx.IndentedJSON(http.StatusOK, connections)
}

func CreateConnection(ctx *gin.Context) {
	connection := serializers.ConnectionCreateSerializer{}

	if err := ctx.BindJSON(&connection); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	if !connection.Valid() {
		ctx.IndentedJSON(http.StatusBadRequest, connection.ValidatorErrs)
		return
	}

	connection_model, err := connection.CreateModel()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	if result := utils.DB.Create(connection_model); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

// todo number of days
func AssignToConnection(ctx *gin.Context) {
	id := ctx.Param("id")
	connection_model := models.Connection{}
	res := utils.DB.First(&connection_model, "id=?", id)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
		return
	}
	connection := serializers.ConnectionAssignSerializer{}
	if err := ctx.BindJSON(&connection); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	connection.DepartureTime = connection_model.DepartureTime.Format("2006-01-02 15:04")
	connection.ArrivalTime = connection_model.ArrivalTime
	if !connection.Valid(int(connection_model.ID)) {
		ctx.IndentedJSON(http.StatusBadRequest, connection.ValidatorErrs)
		return
	}
	connection_model.VehicleRegistration = connection.VehicleReg
	connection_model.DriverID = connection.DriverID
	if result := utils.DB.Save(&connection_model); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

func UpdateConnection(ctx *gin.Context) {
	id := ctx.Param("id")
	connection_model := models.Connection{}
	models_to_change := []models.Connection{}
	res := utils.DB.First(&connection_model, "id=?", id)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
		return
	}
	connection := serializers.ConnectionUpdateSerializer{}
	if err := ctx.BindJSON(&connection); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	//todo asi funguje ???? xd
	orig_deptime := connection_model.DepartureTime
	for i := 0; i < connection.NumberOfDays; i++ {
		connection.ArrivalTime = connection_model.ArrivalTime
		if connection.LineName != connection_model.LineName {
			validators.Line_name_validator(connection.LineName, &connection.ValidatorErrs)
			connection_model.LineName = connection.LineName
		}
		if connection.VehicleReg != connection_model.VehicleRegistration {
			validators.Vehicle_registration_validator(connection.VehicleReg, &connection.ValidatorErrs)
			connection_model.VehicleRegistration = connection.VehicleReg
		}
		if connection.DriverID != connection_model.DriverID {
			validators.Driver_id_validator(connection.DriverID, &connection.ValidatorErrs)
			connection_model.DriverID = connection.DriverID
		}
		if len(connection.ValidatorErrs) != 0 {
			ctx.IndentedJSON(http.StatusBadRequest, connection.ValidatorErrs)
			return
		}
		dep_time, err := time.Parse("2006-01-02 15:04", connection.DepartureTime)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		if !connection_model.DepartureTime.Equal(dep_time) {
			arr_time := serializers.Get_arrival_time(dep_time, connection_model.LineName)
			validators.Driver_availability(int(connection_model.ID), connection_model.DriverID, connection.DepartureTime, arr_time, 1, &connection.ValidatorErrs)
			validators.Vehicle_availability(int(connection_model.ID), connection_model.VehicleRegistration, connection.DepartureTime, arr_time, 1, &connection.ValidatorErrs)
			connection_model.DepartureTime = dep_time
			connection_model.ArrivalTime = arr_time
		}
		if len(connection.ValidatorErrs) != 0 {
			ctx.IndentedJSON(http.StatusBadRequest, connection.ValidatorErrs)
			return
		}
		models_to_change = append(models_to_change, connection_model)
		connection_model = models.Connection{}
		orig_deptime = orig_deptime.AddDate(0, 0, 1)
		res := utils.DB.Where("departure_time=? AND line_name=?", orig_deptime, connection.LineName).Find(&connection_model)
		if res.Error != nil {
			ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
			return
		}
		if res.RowsAffected == 0 {
			break
		}
		connection.DepartureTime = connection_model.DepartureTime.Format("2006-01-02 15:04")
		orig_deptime = connection_model.DepartureTime

	}
	for i := 0; i < len(models_to_change); i++ {
		res := utils.DB.Where("departure_time=? AND line_name=?", models_to_change[i].DepartureTime, models_to_change[i].LineName).Find(&models.Connection{})
		if res.Error != nil {
			ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
			return
		}
		if res.RowsAffected != 0 {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Some connection at same time for this line already exists"})
			return
		}
	}
	for i := 0; i < len(models_to_change); i++ {
		fmt.Print("Models to change: ", models_to_change[i].ID, "\n") //todo debug print
		if result := utils.DB.Save(&models_to_change[i]); result.Error != nil {
			ctx.IndentedJSON(http.StatusBadRequest, result.Error)
			return
		} else {
			ctx.IndentedJSON(http.StatusOK, result)
		}
	}
}

func DeleteConnection(ctx *gin.Context) {
	id := ctx.Param("id")
	connection_model := models.Connection{}
	res := utils.DB.First(&connection_model, "id=?", id)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
		return
	}
	if result := utils.DB.Delete(&connection_model); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Connection deleted successfully"})
	}
}

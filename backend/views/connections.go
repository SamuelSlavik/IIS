package views

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/serializers"
	"github.com/AdamPekny/IIS/backend/utils"
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
		connection := serializers.ConnectionSerializer{
			ID:       model.ID,
			LineName: model.LineName,
		}
		var vehicle models.Vehicle
		err = utils.DB.First(&vehicle, "id=?", model.VehicleID).Error
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		connection.Type = vehicle.VehicleTypeName
		connection.ListStops, err = getStops(connection.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		connections = append(connections, connection)

	}
	ctx.IndentedJSON(http.StatusOK, connections)
}

func GetConnection(ctx *gin.Context) {
	id := ctx.Param("id")
	var connection_model models.Connection
	var err error
	err = utils.DB.First(&connection_model, "id=?", id).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	connection := serializers.ConnectionSerializer{
		ID:       connection_model.ID,
		LineName: connection_model.LineName,
	}
	var vehicle models.Vehicle
	err = utils.DB.First(&vehicle, "id=?", connection_model.VehicleID).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	connection.Type = vehicle.VehicleTypeName
	connection.ListStops, err = getStops(connection.ID)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, connection)
}
func getStops(id uint) (*[]serializers.StopInConnection, error) {
	var connection models.Connection
	stops := []serializers.StopInConnection{}
	if result := utils.DB.First(&connection, id); result.Error != nil {
		return nil, result.Error
	} else {
		var line models.Line
		utils.DB.Model(&line).Preload("Segments").First(&line, "Name = ?", connection.LineName) //todo error handling
		stop1 := line.InitialStop
		dep_time := connection.DepartureTime
		for {
			var segment models.Segment
			err := utils.DB.Preload("Stop1").Preload("Stop2").Joins("inner join line_segments on stop_name1=segment_stop_name1 AND stop_name2=segment_stop_name2").
				First(&segment, "stop_name1=? AND line_segments.line_name=?", stop1, line.Name).Error
			if err != nil {
				return nil, err
			}
			if segment.StopName2 == line.FinalStop {
				stops = append(stops, serializers.StopInConnection{
					StopName:      segment.StopName1,
					DepartureTime: dep_time.Format("15:04"),
				})
				dep_time = dep_time.Add(time.Minute * time.Duration(segment.Time))
				stops = append(stops, serializers.StopInConnection{
					StopName:      segment.StopName2,
					DepartureTime: dep_time.Format("15:04"),
				})
				break
			}
			stops = append(stops, serializers.StopInConnection{
				StopName:      segment.StopName1,
				DepartureTime: dep_time.Format("15:04"),
			})
			dep_time = dep_time.Add(time.Minute * time.Duration(segment.Time))
			stop1 = segment.StopName2
		}

	}
	return &stops, nil
}

func GetConnectionsByLine(ctx *gin.Context) {
	line := ctx.Param("line")
	var connection_models []models.Connection
	var err error
	err = utils.DB.Find(&connection_models, "line_name=?", line).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	connections := []serializers.ConnectionLineSerializer{}
	for _, model := range connection_models {
		connection := serializers.ConnectionLineSerializer{
			ConnectionID:  model.ID,
			LineName:      model.LineName,
			ArrivalTime:   model.ArrivalTime.Format("2006-01-02 15:04"),
			DepartureTime: model.DepartureTime.Format("2006-01-02 15:04"),
			Dirrection:    model.Dirrection,
		}
		connections = append(connections, connection)
	}
	ctx.IndentedJSON(http.StatusOK, connections)
}

func GetConnectionsByLineAndDate(ctx *gin.Context) {
	line := ctx.Param("line")
	date := ctx.Param("date")
	var connection_models []models.Connection
	var err error
	err = utils.DB.Find(&connection_models, "line_name=? AND departure_time BETWEEN ? AND ? ", line, date, date+" 23:59:59").Error
	fmt.Print(len(connection_models))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	connections := []serializers.ConnectionLineSerializer{}
	for _, model := range connection_models {
		connection := serializers.ConnectionLineSerializer{
			ConnectionID:  model.ID,
			LineName:      model.LineName,
			ArrivalTime:   model.ArrivalTime.Format("2006-01-02 15:04"),
			DepartureTime: model.DepartureTime.Format("2006-01-02 15:04"),
			Dirrection:    model.Dirrection,
		}
		connections = append(connections, connection)
	}
	ctx.IndentedJSON(http.StatusOK, connections)
}

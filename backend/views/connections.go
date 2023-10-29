package views

import (
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
func getStops(id uint) (*[]serializers.StopsSerializer, error) {
	var connection models.Connection
	stops := []serializers.StopsSerializer{}
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
				stops = append(stops, serializers.StopsSerializer{
					StopName:      segment.StopName1,
					DepartureTime: dep_time.Format("15:04"),
				})
				stops = append(stops, serializers.StopsSerializer{
					StopName:      segment.StopName2,
					DepartureTime: dep_time.Format("15:04"),
				})
				break
			}
			stops = append(stops, serializers.StopsSerializer{
				StopName:      segment.StopName1,
				DepartureTime: dep_time.Format("15:04"),
			})
			dep_time = dep_time.Add(time.Minute * time.Duration(segment.Time))
			stop1 = segment.StopName2
		}

	}
	return &stops, nil
}

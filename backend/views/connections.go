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
		err = utils.DB.First(&vehicle, "id=?", model.VehicleRegistration).Error
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

func GetConnectionById(ctx *gin.Context) {
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
	err = utils.DB.First(&vehicle, "id=?", connection_model.VehicleRegistration).Error
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
	//stop1 := line.InitialStop
	dep_time := connection.DepartureTime
	//todo zvalidovat funkcnost
	for i := 0; i < len(line.Segments); i++ {
		/*var segment models.Segment
		err := utils.DB.First(&segment, "stop_name1=? AND line_name=?", stop1, line.Name).Error
		if err != nil {
			return nil, err
		}*/
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
		//stop1 = segment.StopName2

	}
	return &stops, nil
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
	connections := []serializers.ConnectionLineSerializer{}
	for _, model := range connection_models {
		connection := serializers.ConnectionLineSerializer{
			ConnectionID:  model.ID,
			LineName:      model.LineName,
			ArrivalTime:   model.ArrivalTime.Format("2006-01-02 15:04"),
			DepartureTime: model.DepartureTime.Format("2006-01-02 15:04"),
			Dirrection:    model.Dirrection,
			InitialStop:   line_model.InitialStop,
			FinalStop:     line_model.FinalStop,
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
			InitialStop:   line_model.InitialStop,
			FinalStop:     line_model.FinalStop,
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

	connection_model := connection.CreateModel()

	if result := utils.DB.Create(connection_model); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

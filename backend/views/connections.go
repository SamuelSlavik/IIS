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
	var connections []serializers.ConnectionListSerializer
	utils.DB.Model(&models.Connection{}).Find(&connections)
	ctx.IndentedJSON(http.StatusOK, connections)
}

func GetConnection(ctx *gin.Context) {
	id := ctx.Param("id")
	var connection models.Connection
	if result := utils.DB.First(&connection, id); result.Error != nil {
		fmt.Print(result.Error)
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	} else {
		var line models.Line
		stops := []serializers.ShowConnectionSerializer{}
		utils.DB.Model(&line).Preload("Segments").First(&line, "Name = ?", connection.LineName) //todo error handling
		stop1 := line.InitialStop
		dep_time := connection.DepartureTime
		for {
			var segment models.Segment
			err := utils.DB.Preload("Stop1").Preload("Stop2").Joins("inner join line_segments on stop_name1=segment_stop_name1 AND stop_name2=segment_stop_name2").
				First(&segment, "stop_name1=? AND line_segments.line_name=?", stop1, line.Name).Error
			if err != nil {
				ctx.IndentedJSON(http.StatusBadRequest, err.Error())
				return
			}
			if segment.StopName2 == line.FinalStop {
				stops = append(stops, serializers.ShowConnectionSerializer{
					StopName:      segment.StopName1,
					DepartureTime: dep_time.Format("15:04"),
				})
				stops = append(stops, serializers.ShowConnectionSerializer{
					StopName:      segment.StopName2,
					DepartureTime: dep_time.Format("15:04"),
				})
				break
			}
			stops = append(stops, serializers.ShowConnectionSerializer{
				StopName:      segment.StopName1,
				DepartureTime: dep_time.Format("15:04"),
			})
			dep_time = dep_time.Add(time.Minute * time.Duration(segment.Time))
			stop1 = segment.StopName2
		}
		ctx.IndentedJSON(http.StatusOK, stops)

	}
}

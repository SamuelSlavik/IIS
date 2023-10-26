package views

import (
	"fmt"
	"net/http"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/serializers"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/gin-gonic/gin"
)

func Get_connections(ctx *gin.Context) {
	var connections []serializers.ConnectionListSerializer
	utils.DB.Model(&models.Connection{}).Find(&connections)
	ctx.IndentedJSON(http.StatusOK, connections)
}

func Get_connection(ctx *gin.Context) {
	id := ctx.Param("id")
	var connection models.Connection
	if result := utils.DB.First(&connection, id); result.Error != nil {
		fmt.Print(result.Error)
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	} else {
		var line models.Line
		//stopm := map[models.Stop]bool{}
		stops := []models.Stop{}
		utils.DB.Model(&line).Preload("Segments").First(&line, "Name = ?", connection.LineName) //todo error handling
		for _, segment := range line.Segments {
			utils.DB.Model(&segment).Preload("Stop1").Preload("Stop2").First(&segment)
			stops = append(stops, segment.Stop1)
		}
		fmt.Print(stops, "\n")
		for _, segment := range line.Segments {
			fmt.Print(segment, "\n")
		}
		ctx.IndentedJSON(http.StatusOK, stops)

	}
}

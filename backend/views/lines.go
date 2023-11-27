package views

import (
	"net/http"
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/serializers"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/gin-gonic/gin"
)

func ListLines(ctx *gin.Context) {
	var lineModels []models.Line
	var lineSerializers []serializers.LineInList

	// Fetch all lines from the database
	res := utils.DB.Find(&lineModels)

	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
		return
	}

	// Populate lineSerializers with data from lineModels
	for _, lineModel := range lineModels {
		lineSerializer := serializers.LineInList{}
		lineSerializer.FromModel(lineModel)
		lineSerializers = append(lineSerializers, lineSerializer)
	}

	ctx.IndentedJSON(http.StatusOK, lineSerializers)
}

func GetLine(ctx *gin.Context) {
	lineName := ctx.Param("line")
	var lineSerializer serializers.LineSerializer
	lineSerializer.Name = lineName
	err := lineSerializer.GetStops(lineName)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, lineSerializer)
}

func CreateLine(ctx *gin.Context) {
	var lineSerializer serializers.LineCreateSerializer
	err := ctx.BindJSON(&lineSerializer)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(lineSerializer.StopsSequence) < 2 {
		ctx.IndentedJSON(http.StatusBadRequest, "Line must have at least 2 stops")
		return
	}
	segments := []models.Segment{}
	for i := 0; i < len(lineSerializer.StopsSequence)-1; i++ {
		segment := models.Segment{
			StopName1: lineSerializer.StopsSequence[i].StopName,
			StopName2: lineSerializer.StopsSequence[i+1].StopName,
			Time:      lineSerializer.StopsSequence[i].Duration,
			LineName:  lineSerializer.Name,
		}
		segments = append(segments, segment)
	}
	line := models.Line{
		Name:        lineSerializer.Name,
		InitialStop: lineSerializer.StopsSequence[0].StopName,
		FinalStop:   lineSerializer.StopsSequence[len(lineSerializer.StopsSequence)-1].StopName,
		Segments:    segments,
	}
	res := utils.DB.Create(&line)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, lineSerializer)
}

func UpdateLine(ctx *gin.Context) {
	lineName := ctx.Param("line")
	var lineSerializer serializers.LineUpdateSerializer
	err := ctx.BindJSON(&lineSerializer)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(lineSerializer.StopsSequence) < 2 {
		ctx.IndentedJSON(http.StatusBadRequest, "Line must have at least 2 stops")
		return
	}
	var line models.Line
	res := utils.DB.Preload("Segments").First(&line, "Name = ?", lineName)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
		return
	}
	for _, segment := range line.Segments {
		res := utils.DB.Delete(&segment)
		if res.Error != nil {
			ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
			return
		}
	}
	segments := []models.Segment{}
	var duration uint
	for i := 0; i < len(lineSerializer.StopsSequence)-1; i++ {
		duration += lineSerializer.StopsSequence[i].Duration
		segment := models.Segment{
			StopName1: lineSerializer.StopsSequence[i].StopName,
			StopName2: lineSerializer.StopsSequence[i+1].StopName,
			Time:      lineSerializer.StopsSequence[i].Duration,
		}
		segments = append(segments, segment)
	}
	line.InitialStop = lineSerializer.StopsSequence[0].StopName
	line.FinalStop = lineSerializer.StopsSequence[len(lineSerializer.StopsSequence)-1].StopName
	line.Segments = segments
	res = utils.DB.Save(&line)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
		return
	}
	connections := []models.Connection{}
	res = utils.DB.Find(&connections, "line_name = ?", lineName)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
		return
	}

	for _, connection := range connections {
		connection.ArrivalTime = connection.DepartureTime.Add(time.Minute * time.Duration(duration))
		res = utils.DB.Save(&connection)
		if res.Error != nil {
			ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
			return
		}

	}
	ctx.IndentedJSON(http.StatusOK, lineSerializer)
}

func DeleteLine(ctx *gin.Context) {
	lineName := ctx.Param("line")
	res := utils.DB.Delete(&models.Line{Name: lineName})
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Line deleted successfully"})
}

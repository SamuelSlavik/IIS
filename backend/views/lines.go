package views

import (
	"net/http"

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
	res := utils.DB.Create(&models.Line{Name: lineSerializer.Name})
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error.Error())
		return
	}
	for i := 0; i < len(lineSerializer.StopsSequence)-1; i++ {
		segment := models.Segment{
			StopName1: lineSerializer.StopsSequence[i].StopName,
			StopName2: lineSerializer.StopsSequence[i+1].StopName,
			Time:      lineSerializer.StopsSequence[i].Duration,
			LineName:  lineSerializer.Name,
		}
		res := utils.DB.Create(&segment)
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
	ctx.IndentedJSON(http.StatusOK, "ok")
}

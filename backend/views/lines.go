package views

import (
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/serializers"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListLines(ctx *gin.Context) {
	var lineModels []models.Line
	var lineSerializers []serializers.LineInList

	// Fetch all lines from the database
	res := utils.DB.Find(&lineModels)

	if res.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Could not retrieve lines",
		})
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

}

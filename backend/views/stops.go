package views

import (
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/serializers"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ListStops(ctx *gin.Context) {
	var stopModels []models.Stop
	var stopSerializers []serializers.StopSerializer

	// Retrieve the query parameter from the URL
	query := ctx.Query("query")

	// Build the query to filter stops by name if the query parameter is provided
	dbQuery := utils.DB.Order("name")
	if query != "" {
		dbQuery = dbQuery.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(query)+"%")
	}

	// Fetch stops from the database based on the query
	res := dbQuery.Find(&stopModels)

	if res.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Could not retrieve stops!",
		})
		return
	}

	for _, stopModel := range stopModels {
		stopSerializer := serializers.StopSerializer{}
		stopSerializer.FromModel(stopModel)
		stopSerializers = append(stopSerializers, stopSerializer)
	}

	ctx.IndentedJSON(http.StatusOK, stopSerializers)
}

func GetStop(ctx *gin.Context) {
	stopID := ctx.Param("id")

	var stopModel models.Stop
	var stopSerializer serializers.StopSerializer

	res := utils.DB.First(&stopModel, stopID)

	if res.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Stop not found",
		})
		return
	}

	stopSerializer.FromModel(stopModel)

	ctx.IndentedJSON(http.StatusOK, stopSerializer)
}

func EditStop(ctx *gin.Context) {
	stopID := ctx.Param("id")

	var existingStop models.Stop
	result := utils.DB.First(&existingStop, stopID)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Stop not found",
		})
		return
	}

	var editRequest serializers.EditStopSerializer
	if err := ctx.BindJSON(&editRequest); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	existingStop.Name = editRequest.Name
	result = utils.DB.Save(&existingStop)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update stop",
		})
		return
	}

	stopSerializer := serializers.StopSerializer{
		ID:   existingStop.ID,
		Name: existingStop.Name,
	}

	ctx.IndentedJSON(http.StatusOK, stopSerializer)
}

func CreateStop(ctx *gin.Context) {
	// Bind the form data to a StopCreateRequest struct
	var createRequest serializers.StopCreateRequest
	if err := ctx.Bind(&createRequest); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	// Validate the input data
	if createRequest.Name == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
		return
	}

	// Create a new stop
	newStop := models.Stop{
		Name: createRequest.Name,
		// Add other fields if needed
	}

	// Save the new stop to the database
	result := utils.DB.Create(&newStop)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create stop",
		})
		return
	}

	// Return the created stop in the response
	stopSerializer := serializers.StopSerializer{
		ID:   newStop.ID,
		Name: newStop.Name,
		// Add other fields if needed
	}

	ctx.IndentedJSON(http.StatusOK, stopSerializer)
}

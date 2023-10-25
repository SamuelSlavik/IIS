package views

import (
	"net/http"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/serializers"
	utils "github.com/AdamPekny/IIS/backend/utils"
	"github.com/gin-gonic/gin"
)

func List_vehicles(ctx *gin.Context) {
	var vehicles []models.Vehicle
	utils.DB.Preload("VehicleType").Find(&vehicles)
	ctx.IndentedJSON(http.StatusOK, vehicles)
}

func Create_vehicle(ctx *gin.Context) {
	vehicle := serializers.VehicleSerializer{}

	if err := ctx.BindJSON(&vehicle); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	if !vehicle.Valid() {
		ctx.IndentedJSON(http.StatusBadRequest, vehicle.ValidatorErrs)
		return
	}

	vehicle_model := vehicle.Create_model()

	if result := utils.DB.Create(vehicle_model); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

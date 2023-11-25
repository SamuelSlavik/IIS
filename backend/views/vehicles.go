package views

import (
	"fmt"
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

func GetVehicle(ctx *gin.Context) {
	vehicle_id := ctx.Param("id")
	vehicle := models.Vehicle{}
	mainteneces := []models.MaintenanceRequest{}
	res := utils.DB.Preload("VehicleType").First(&vehicle, "registration = ?", vehicle_id)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error)
		return
	}
	res = utils.DB.Joins("JOIN malfunction_reports ON maintenance_requests.malfunc_rep_ref = malfunction_reports.id").Where("vehicle_ref = ?", vehicle_id).
		Order("created_at DESC").Find(&mainteneces)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error)
		return
	}
	fmt.Print(mainteneces[0].Status, "\n")
	if mainteneces[0].Status == "done" {
		vehicle_serializer := serializers.VehicleGetSerializer{
			Registration: vehicle.Registration,
			Capacity:     vehicle.Capacity,
			Brand:        vehicle.Brand,
			Type:         vehicle.VehicleType.Type,
			LastMaintenance: serializers.LastMaintenance{
				Status: string(mainteneces[0].Status),
				Date:   mainteneces[0].MaintenRep.CreatedAt.Format("2006-01-02 15:04:05"),
			},
		}
		ctx.IndentedJSON(http.StatusOK, vehicle_serializer)
		return
	} else {
		vehicle_serializer := serializers.VehicleGetSerializer{
			Registration: vehicle.Registration,
			Capacity:     vehicle.Capacity,
			Brand:        vehicle.Brand,
			Type:         vehicle.VehicleType.Type,
			LastMaintenance: serializers.LastMaintenance{
				Status: string(mainteneces[0].Status),
				Date:   mainteneces[0].CreatedAt.Format("2006-01-02 "),
			},
		}
		ctx.IndentedJSON(http.StatusOK, vehicle_serializer)
		return
	}
}

func UpdateVehicle(ctx *gin.Context) {
	vehicle_id := ctx.Param("id")
	vehicle := models.Vehicle{}
	res := utils.DB.First(&vehicle, "registration = ?", vehicle_id)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error)
		return
	}
	vehicle_serializer := serializers.VehicleUpdateSerializer{}
	if err := ctx.BindJSON(&vehicle_serializer); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if !vehicle_serializer.Valid() {
		ctx.IndentedJSON(http.StatusBadRequest, vehicle_serializer.ValidatorErrs)
		return
	}
	vehicle.Capacity = vehicle_serializer.Capacity
	if vehicle_serializer.Brand != "" {
		vehicle.Brand = vehicle_serializer.Brand
	}
	vehicle.VehicleTypeName = vehicle_serializer.Type
	if result := utils.DB.Save(&vehicle); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

func DeleteVehicle(ctx *gin.Context) {
	vehicle_id := ctx.Param("id")
	vehicle := models.Vehicle{}
	res := utils.DB.First(&vehicle, "registration = ?", vehicle_id)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error)
		return
	}
	if result := utils.DB.Delete(&vehicle); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

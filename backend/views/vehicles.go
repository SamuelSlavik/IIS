// package views contains views used in router handlers
package views

import (
	"net/http"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/serializers"
	utils "github.com/AdamPekny/IIS/backend/utils"
	"github.com/gin-gonic/gin"
)

// List_vehicles lists every vehicle in the database
func List_vehicles(ctx *gin.Context) {
	var vehicles []models.Vehicle
	res := utils.DB.Preload("VehicleType").Find(&vehicles)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error)
		return
	}
	var vehicle_serializers []serializers.VehicleGetSerializer
	for _, vehicle := range vehicles {
		mainteneces := []models.MaintenanceRequest{}
		res := utils.DB.Joins("JOIN malfunction_reports ON maintenance_requests.malfunc_rep_ref = malfunction_reports.id").Where("vehicle_ref = ?", vehicle.Registration).
			Order("created_at DESC").Find(&mainteneces)
		if res.Error != nil {
			ctx.IndentedJSON(http.StatusBadRequest, res.Error)
			return
		}
		if len(mainteneces) == 0 {
			vehicle_serializer := serializers.VehicleGetSerializer{
				Registration: vehicle.Registration,
				Capacity:     vehicle.Capacity,
				Brand:        vehicle.Brand,
				Type:         vehicle.VehicleType.Type,
				LastMaintenance: serializers.LastMaintenance{
					Status: "-",
					Date:   "-",
				},
			}
			vehicle_serializers = append(vehicle_serializers, vehicle_serializer)
			continue
		}
		if mainteneces[0].Status == "done" {
			vehicle_serializer := serializers.VehicleGetSerializer{
				Registration: vehicle.Registration,
				Capacity:     vehicle.Capacity,
				Brand:        vehicle.Brand,
				Type:         vehicle.VehicleType.Type,
				LastMaintenance: serializers.LastMaintenance{
					Status: string(mainteneces[0].Status),
					Date:   mainteneces[0].MaintenRep.CreatedAt.Format("2006-01-02 15:04"),
				},
			}
			vehicle_serializers = append(vehicle_serializers, vehicle_serializer)
		} else {
			vehicle_serializer := serializers.VehicleGetSerializer{
				Registration: vehicle.Registration,
				Capacity:     vehicle.Capacity,
				Brand:        vehicle.Brand,
				Type:         vehicle.VehicleType.Type,
				LastMaintenance: serializers.LastMaintenance{
					Status: string(mainteneces[0].Status),
					Date:   mainteneces[0].CreatedAt.Format("2006-01-02"),
				},
			}
			vehicle_serializers = append(vehicle_serializers, vehicle_serializer)
		}
	}
	ctx.IndentedJSON(http.StatusOK, vehicle_serializers)
}

// Create_vehicle creates a new vehicle
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

// GetVehicle gets a vehicle by its registration number
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
	if len(mainteneces) == 0 {
		vehicle_serializer := serializers.VehicleGetSerializer{
			Registration: vehicle.Registration,
			Capacity:     vehicle.Capacity,
			Brand:        vehicle.Brand,
			Type:         vehicle.VehicleType.Type,
			LastMaintenance: serializers.LastMaintenance{
				Status: "-",
				Date:   "-",
			},
		}
		ctx.IndentedJSON(http.StatusOK, vehicle_serializer)
		return
	}
	if mainteneces[0].Status == "done" {
		vehicle_serializer := serializers.VehicleGetSerializer{
			Registration: vehicle.Registration,
			Capacity:     vehicle.Capacity,
			Brand:        vehicle.Brand,
			Type:         vehicle.VehicleType.Type,
			LastMaintenance: serializers.LastMaintenance{
				Status: string(mainteneces[0].Status),
				Date:   mainteneces[0].MaintenRep.CreatedAt.Format("2006-01-02 15:04"),
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
				Date:   mainteneces[0].CreatedAt.Format("2006-01-02"),
			},
		}
		ctx.IndentedJSON(http.StatusOK, vehicle_serializer)
		return
	}
}

// UpdateVehicle updates a vehicle with given attributes
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

// DeleteVehicle deletes a vehicle
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
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Vehicle deleted successfully"})
	}
}

// ListNotBrokenVehicles lists vehicles that does not have any malfunction report or all of its malfunction reports are done
func ListNotBrokenVehicles(ctx *gin.Context) {
	var vehicles []models.Vehicle
	res := utils.DB.Preload("VehicleType").Find(&vehicles)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error)
		return
	}
	result := utils.DB.Table("vehicles").Joins("LEFT JOIN malfunction_reports ON vehicles.registration = malfunction_reports.vehicle_ref").
		Joins("LEFT JOIN maintenance_requests ON maintenance_requests.malfunc_rep_ref = malfunction_reports.id").
		Group("vehicles.registration").
		Having("COUNT(DISTINCT malfunction_reports.id) = SUM(CASE WHEN maintenance_requests.status = 'done' THEN 1 ELSE 0 END)").Preload("VehicleType").
		Find(&vehicles)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}
	var vehicle_serializers []serializers.VehicleGetSerializer
	for _, vehicle := range vehicles {
		mainteneces := []models.MaintenanceRequest{}
		res := utils.DB.Joins("JOIN malfunction_reports ON maintenance_requests.malfunc_rep_ref = malfunction_reports.id").
			Where("vehicle_ref = ?", vehicle.Registration).
			Order("created_at DESC").Find(&mainteneces)
		if res.Error != nil {
			ctx.IndentedJSON(http.StatusBadRequest, res.Error)
			return
		}
		if len(mainteneces) == 0 {
			vehicle_serializer := serializers.VehicleGetSerializer{
				Registration: vehicle.Registration,
				Capacity:     vehicle.Capacity,
				Brand:        vehicle.Brand,
				Type:         vehicle.VehicleType.Type,
				LastMaintenance: serializers.LastMaintenance{
					Status: "-",
					Date:   "-",
				},
			}
			vehicle_serializers = append(vehicle_serializers, vehicle_serializer)
		} else {
			vehicle_serializer := serializers.VehicleGetSerializer{
				Registration: vehicle.Registration,
				Capacity:     vehicle.Capacity,
				Brand:        vehicle.Brand,
				Type:         vehicle.VehicleType.Type,
				LastMaintenance: serializers.LastMaintenance{
					Status: string(mainteneces[0].Status),
					Date:   mainteneces[0].CreatedAt.Format("2006-01-02"),
				},
			}
			vehicle_serializers = append(vehicle_serializers, vehicle_serializer)
		}

	}
	ctx.IndentedJSON(http.StatusOK, vehicle_serializers)
}

// CreateVehicleType creates a new vehicle unique vehicle type
func CreateVehicleType(ctx *gin.Context) {
	vehicle_type := serializers.VehicleTypeCreateSerializer{}

	if err := ctx.BindJSON(&vehicle_type); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	res := utils.DB.Where("type = ?", vehicle_type.Type).Find(&models.VehicleType{})
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, res.Error)
		return
	}
	if res.RowsAffected > 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Vehicle type already exists"})
		return
	}

	vehicle_type_model := models.VehicleType{
		Type: vehicle_type.Type,
	}
	if result := utils.DB.Create(&vehicle_type_model); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

// ListVehicleTypes lists all vehicle types
func ListVehicleTypes(ctx *gin.Context) {
	var vehicle_types []models.VehicleType
	if result := utils.DB.Find(&vehicle_types); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}
	var vehicle_type_serializers []serializers.VehicleTypeSerializer
	for _, vehicle_type := range vehicle_types {
		vehicle_type_serializer := serializers.VehicleTypeSerializer{
			Type: vehicle_type.Type,
		}
		result := utils.DB.Where("vehicle_type_name=?", vehicle_type.Type).Find(&models.Vehicle{})
		if result.Error != nil {
			ctx.IndentedJSON(http.StatusBadRequest, result.Error)
			return
		}
		if result.RowsAffected > 0 {
			vehicle_type_serializer.Active = true
		} else {
			vehicle_type_serializer.Active = false
		}
		vehicle_type_serializers = append(vehicle_type_serializers, vehicle_type_serializer)
	}
	ctx.IndentedJSON(http.StatusOK, vehicle_type_serializers)
}

// DeleteVehicleType deletes a vehicle type if it is not used by any vehicle
func DeleteVehicleType(ctx *gin.Context) {
	vehicle_type := ctx.Param("id")
	var vehicle_type_model models.VehicleType
	if result := utils.DB.First(&vehicle_type_model, "id = ?", vehicle_type); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}
	vehicles := []models.Vehicle{}
	result := utils.DB.Where("vehicle_type_name=?", vehicle_type_model.Type).Find(&vehicles)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}
	if len(vehicles) > 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Vehicle type is still in use"})
		return
	}
	if result := utils.DB.Delete(&vehicle_type_model); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Vehicle type deleted successfully"})
	}
}

package api

import (
	"github.com/AdamPekny/IIS/backend/middleware"
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/views"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:5173"}

	router.Use(cors.New(config))

	// Users
	router.GET("/api/users/list", middleware.RequireAuth(), views.ListUsers)
	router.GET("/api/users/list/:role", middleware.RequireAuth(), views.ListRoleUsers)
	router.GET("/api/users/get/:id", middleware.RequireAuth(), views.RetrieveUser)
	router.GET("/api/users/get", middleware.RequireAuth(), views.RetrieveCurrentUser)
	router.GET("/api/users/logout", middleware.RequireAuth(), views.Logout)

	router.POST("/api/users/signup", views.Signup)
	router.POST("/api/users/login", views.Login)

	router.PATCH("/api/users/update/:id", middleware.RequireAuth(), views.UpdateUser)

	router.DELETE("/api/users/delete/:id", middleware.RequireAuth(string(models.AdminRole)), views.DeleteUser)

	// Vehicles
	router.GET("/api/vehicles/list", middleware.RequireAuth(string(models.SuperuserRole), string(models.DriverRole)), views.List_vehicles)
	router.GET("/api/vehicles/get/:id", middleware.RequireAuth(string(models.SuperuserRole)), views.GetVehicle)
	router.POST("/api/vehicles/create", middleware.RequireAuth(string(models.SuperuserRole)), views.Create_vehicle)
	router.PUT("/api/vehicles/update/:id", middleware.RequireAuth(string(models.SuperuserRole)), views.UpdateVehicle)
	router.DELETE("/api/vehicles/delete/:id", middleware.RequireAuth(string(models.SuperuserRole)), views.DeleteVehicle)
	router.GET("/api/vehicles/list/ok", middleware.RequireAuth(string(models.SuperuserRole), string(models.DispatcherRole)), views.ListNotBrokenVehicles)
	router.GET("/api/vehicletypes/list", middleware.RequireAuth(string(models.SuperuserRole)), views.ListVehicleTypes)
	router.POST("/api/vehicletypes/create", middleware.RequireAuth(string(models.SuperuserRole)), views.CreateVehicleType)
	router.DELETE("/api/vehicletypes/delete/:id", middleware.RequireAuth(string(models.SuperuserRole)), views.DeleteVehicleType)

	// Connections
	router.GET("/api/connections/list", middleware.RequireAuth(string(models.SuperuserRole), string(models.DispatcherRole)), views.ListConnections)
	router.GET("/api/connections/list/:line", middleware.RequireAuth(string(models.SuperuserRole), string(models.DispatcherRole)), views.ListConnectionsByLine)
	router.GET("/api/connections/list/:line/:date", middleware.RequireAuth(string(models.SuperuserRole), string(models.DispatcherRole)), views.ListConnectionsByLineAndDate)
	router.GET("/api/connections/get/:id", middleware.RequireAuth(string(models.SuperuserRole), string(models.DispatcherRole), string(models.DriverRole)), views.GetConnectionById)
	router.POST("/api/connections/create", middleware.RequireAuth(string(models.SuperuserRole)), views.CreateConnection)
	router.PATCH("/api/connections/assign/:id", middleware.RequireAuth(string(models.DispatcherRole)), views.AssignToConnection)
	router.PATCH("/api/connections/update/:id", middleware.RequireAuth(string(models.SuperuserRole)), views.UpdateConnection)
	router.DELETE("/api/connections/delete/:id/:days", middleware.RequireAuth(string(models.SuperuserRole)), views.DeleteConnection)
	router.GET("/api/connections/list/driver/:id", middleware.RequireAuth(string(models.DriverRole)), views.ListDriverConnections)
	// not logged user
	router.GET("/api/connections/search", views.ListUserConnections)
	router.GET("/api/connections/search/:line", views.ListUserConnectionsByLine)
	router.GET("/api/connections/search/:line/:date", views.ListConnectionsUserByLineAndDate)
	router.GET("/api/connections/get/details/:id", views.GetDetailOfConnection) //unregistered ???
	//stops
	router.GET("/api/stops", middleware.RequireAuth(string(models.SuperuserRole)), views.ListStops)
	router.GET("/api/stops/get/:id", middleware.RequireAuth(string(models.SuperuserRole)), views.GetStop)
	router.POST("/api/stops/create", middleware.RequireAuth(string(models.SuperuserRole)), views.CreateStop)
	router.PUT("/api/stops/edit/:id", middleware.RequireAuth(string(models.SuperuserRole)), views.EditStop)
	router.DELETE("/api/stops/delete/:id", middleware.RequireAuth(string(models.SuperuserRole)), views.DeleteStop)

	//lines
	router.GET("/api/lines/list", middleware.RequireAuth(string(models.SuperuserRole)), views.ListLines)
	router.GET("/api/lines/get/:line", middleware.RequireAuth(string(models.SuperuserRole)), views.GetLine)
	router.POST("/api/lines/create", middleware.RequireAuth(string(models.SuperuserRole)), views.CreateLine)
	router.PATCH("/api/lines/update/:line", middleware.RequireAuth(string(models.SuperuserRole)), views.UpdateLine)
	router.DELETE("/api/lines/delete/:line", middleware.RequireAuth(string(models.SuperuserRole)), views.DeleteLine)

	// Maintenance
	router.POST("/api/maintenance/malfunc/create", middleware.RequireAuth(string(models.DriverRole)), views.CreateMalfuncReport)
	router.POST("/api/maintenance/maintenreq/create", middleware.RequireAuth(string(models.SuperuserRole)), views.CreateMaintenRequest)
	router.POST("/api/maintenance/maintenrep/create", middleware.RequireAuth(string(models.TechnicianRole)), views.CreateMaintenReport)
	router.PUT("/api/maintenance/malfunc/update/:id", middleware.RequireAuth(string(models.DriverRole)), views.UpdateMalfuncReport)
	router.PUT("/api/maintenance/maintenreq/update/:id", middleware.RequireAuth(string(models.SuperuserRole)), views.UpdateMaintenRequest)
	router.PATCH("/api/maintenance/maintenreq/assigntech/:id", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.AssignTechMaintenRequest)
	router.PUT("/api/maintenance/maintenrep/update/:id", middleware.RequireAuth(string(models.TechnicianRole)), views.UpdateMaintenReport)
	router.DELETE("/api/maintenance/malfunc/delete/:id", middleware.RequireAuth(string(models.DriverRole)), views.DeleteMalfuncReport)
	router.DELETE("/api/maintenance/maintenreq/delete/:id", middleware.RequireAuth(string(models.DriverRole)), views.DeleteMaintenRequest)
	router.DELETE("/api/maintenance/maintenrep/delete/:id", middleware.RequireAuth(string(models.TechnicianRole)), views.DeleteMaintenReport)
	router.GET("/api/maintenance/malfunc/list", middleware.RequireAuth(string(models.DriverRole), string(models.SuperuserRole)), views.ListStatusMalfuncReports)
	router.GET("/api/maintenance/malfunc/get/:id", middleware.RequireAuth(string(models.DriverRole), string(models.SuperuserRole)), views.GetMalfuncReport)
	router.GET("/api/maintenance/maintenreq/list", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListStatusMaintenRequests)
	router.GET("/api/maintenance/maintenreq/list/unassigned", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListUnassignedMaintenRequests)
	router.GET("/api/maintenance/maintenreq/list/super/:id/:status", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListCreatorStatusMaintenRequests)
	router.GET("/api/maintenance/maintenreq/list/tech/:id/:status", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListResolverStatusMaintenRequests)
	router.GET("/api/maintenance/maintenreq/get/:id", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.GetMaintenRequest)
	router.GET("/api/maintenance/maintenrep/list", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListMaintenReports)
	router.GET("/api/maintenance/maintenrep/get/:id", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.GetMaintenReport)

	return router
}

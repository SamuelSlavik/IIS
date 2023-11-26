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
	router.GET("/api/users/get/:id", middleware.RequireAuth(), views.RetrieveUser)
	router.GET("/api/users/get", middleware.RequireAuth(), views.RetrieveCurrentUser)
	router.GET("/api/users/logout", middleware.RequireAuth(), views.Logout)

	router.POST("/api/users/signup", views.Signup)
	router.POST("/api/users/login", views.Login)

	router.PATCH("/api/users/update/:id", middleware.RequireAuth(), views.UpdateUser)

	router.DELETE("/api/users/delete/:id", middleware.RequireAuth(string(models.AdminRole)), views.DeleteUser)

	// Vehicles
	router.GET("/api/vehicles/list", views.List_vehicles)
	router.GET("/api/vehicles/get/:id", views.GetVehicle)
	router.POST("/api/vehicles/create", views.Create_vehicle)
	router.PUT("/api/vehicles/update/:id", views.UpdateVehicle)
	router.DELETE("/api/vehicles/delete/:id", views.DeleteVehicle)

	// Connections
	router.GET("/api/connections/list", views.ListConnections)
	router.GET("/api/connections/list/:line", views.ListConnectionsByLine)
	router.GET("/api/connections/list/:line/:date", views.ListConnectionsByLineAndDate)
	router.GET("/api/connections/get/:id", views.GetConnectionById)
	router.GET("/api/connections/get/details/:id", views.GetDetailOfConnection)
	router.POST("/api/connections/create", views.CreateConnection)
	router.PATCH("/api/connections/assign/:id", views.AssignToConnection)
	router.PATCH("/api/connections/update/:id", views.UpdateConnection)
	router.DELETE("/api/connections/delete/:id", views.DeleteConnection)

	router.GET("/api/stops", views.ListStops)
	router.GET("/api/stops/get/:id", views.GetStop)
	//router.DELETE("/api/stops/delete/:id", views.DeleteStop)
	router.PUT("/api/stops/edit/:id", views.EditStop)
	router.POST("/api/stops/create", views.CreateStop)

	router.GET("/api/lines", views.ListLines)
	router.GET("/api/lines/get/:id", views.GetLine)
	//router.POST("/api/lines/create", views.CreateLine)
	//router.PUT("/api/lines/edit/:id", views.EditLine)
	//router.DELETE("/api/lines/delete/:id", views.DeleteLine)

	// Maintenance
	router.POST("/api/maintenance/malfunc/create", middleware.RequireAuth(string(models.DriverRole)), views.CreateMalfuncReport)
	router.POST("/api/maintenance/maintenreq/create", middleware.RequireAuth(string(models.SuperuserRole)), views.CreateMaintenRequest)
	router.POST("/api/maintenance/maintenrep/create", middleware.RequireAuth(string(models.TechnicianRole)), views.CreateMaintenReport)
	router.PUT("/api/maintenance/malfunc/update/:id", middleware.RequireAuth(string(models.DriverRole)), views.UpdateMalfuncReport)
	router.PUT("/api/maintenance/maintenreq/update/:id", middleware.RequireAuth(string(models.SuperuserRole)), views.UpdateMaintenRequest)
	router.DELETE("/api/maintenance/malfunc/delete/:id", middleware.RequireAuth(string(models.DriverRole)), views.DeleteMalfuncReport)
	router.DELETE("/api/maintenance/maintenreq/delete/:id", middleware.RequireAuth(string(models.DriverRole)), views.DeleteMalfuncReport)
	router.GET("/api/maintenance/malfunc/list", middleware.RequireAuth(string(models.DriverRole), string(models.SuperuserRole)), views.ListStatusMalfuncReports)
	router.GET("/api/maintenance/malfunc/get/:id", middleware.RequireAuth(string(models.DriverRole), string(models.SuperuserRole)), views.GetMalfuncReport)
	router.GET("/api/maintenance/maintenreq/list", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListStatusMaintenRequests)
	router.GET("/api/maintenance/maintenreq/list/unassigned", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListUnassignedMaintenRequests)
	router.GET("/api/maintenance/maintenreq/list/super/:id/:status", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListCreatorStatusMaintenRequests)
	router.GET("/api/maintenance/maintenreq/list/tech/:id/:status", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListResolverStatusMaintenRequests)
	router.GET("/api/maintenance/maintenreq/get/:id", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.GetMaintenRequest)
	router.GET("/api/maintenance/maintenrep/list", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListMaintenReports)

	return router
}

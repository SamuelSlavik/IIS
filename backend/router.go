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
	config.AllowOrigins = []string{"http://127.0.0.1:5173"}

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

	router.POST("/api/vehicles/create", views.Create_vehicle)

	// Connections
	router.GET("/api/connections", views.ListConnections)
	router.GET("/api/connections/get/:id", views.GetConnectionById)
	router.GET("/api/connections/list/:line", views.ListConnectionsByLine)
	router.GET("/api/connections/list/:line/:date", views.ListConnectionsByLineAndDate)
	router.POST("/api/connections/create", views.CreateConnection)
	router.PATCH("/api/connections/assign/:id", views.AssignToConnection)

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
	router.GET("/api/maintenance/maintenreq/list", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListMaintenRequests)
	router.GET("/api/maintenance/maintenreq/list/:status", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListStatusMaintenRequests)
	router.GET("/api/maintenance/maintenreq/list/super/:id/:status", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListCreatorStatusMaintenRequests)
	router.GET("/api/maintenance/maintenreq/list/tech/:id/:status", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListResolverStatusMaintenRequests)
	router.GET("/api/maintenance/maintenreq/get/:id", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.GetMaintenRequest)
	router.GET("/api/maintenance/maintenrep/list", middleware.RequireAuth(string(models.TechnicianRole), string(models.SuperuserRole)), views.ListMaintenReports)

	return router
}

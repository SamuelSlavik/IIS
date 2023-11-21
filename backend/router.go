package api

import (
	"github.com/AdamPekny/IIS/backend/middleware"
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

	router.GET("/api/users/list", middleware.RequireAuth(), views.ListUsers)
	router.GET("/api/users/get/:id", middleware.RequireAuth(), views.RetrieveUser)
	router.GET("/api/users/get", middleware.RequireAuth(), views.RetrieveCurrentUser)

	router.POST("/api/users/signup", views.Signup)
	router.DELETE("/api/users/delete/:id", middleware.RequireAuth(), views.DeleteUser)
	router.POST("/api/users/login", views.Login)
	router.GET("/api/users/logout", views.Logout)

	router.GET("/api/vehicles/list", views.List_vehicles)

	router.POST("/api/vehicles/create", views.Create_vehicle)

	router.GET("/api/connections", views.ListConnections)
	router.GET("/api/connections/:id", views.GetConnection)
	return router
}

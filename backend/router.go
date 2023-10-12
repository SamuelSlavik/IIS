package api

import (
	"net/http"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/api/users/list", func(ctx *gin.Context) {
		var users []models.User
		db, _ := utils.Conn()
		db.Find(&users)
		ctx.IndentedJSON(http.StatusOK, users)
	})

	router.POST("/api/users/create", func(ctx *gin.Context) {
		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.IndentedJSON(http.StatusTeapot, "Error")
			return
		}
		db, _ := utils.Conn()
		result := db.Create(&user)
		ctx.IndentedJSON(http.StatusOK, result)
	})

	return router
}

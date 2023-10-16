package api

import (
	"net/http"

	"github.com/AdamPekny/IIS/backend/userauth"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/api/users/list", func(ctx *gin.Context) {
		var users []userauth.User
		db, _ := utils.Conn()
		db.Preload("UserType").Find(&users)
		ctx.IndentedJSON(http.StatusOK, users)
	})

	router.POST("/api/users/create", userauth.Create_user)

	router.POST("/api/users/type/create", userauth.Create_user_type)

	return router
}

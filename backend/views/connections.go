package views

import (
	"net/http"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/serializers"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/gin-gonic/gin"
)

func Get_connections(ctx *gin.Context) {
	var connections []serializers.ConnectionListSerializer
	utils.DB.Model(&models.Connection{}).Find(&connections)
	ctx.IndentedJSON(http.StatusOK, connections)
}

package views

import (
	"net/http"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/serializers"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/gin-gonic/gin"
)


func CreateMalfuncReport(ctx *gin.Context) {
	var malfunc_report_serializer serializers.MalfuncRepCreateSerialzier

	if err := ctx.BindJSON(&malfunc_report_serializer); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var logged_user_id uint

	if logged_user, ok := ctx.Get("user"); ok {
		logged_user_model, ok := logged_user.(models.User)

		if !ok {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"Error": "Not a valid user!",
			})
			return
		}

		logged_user_id = logged_user_model.ID
	} else {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error": "Not logged in",
		})
		return
	}

	malfunc_report_serializer.CreatedByRef = logged_user_id

	malfunc_report_model := malfunc_report_serializer.ToModel()

	result := utils.DB.Create(&malfunc_report_model)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, malfunc_report_model)
}

func ListMalfuncReports(ctx *gin.Context) {
	var malfunc_reports []models.MalfunctionReport

	if result := utils.DB.Preload("Creator").Find(&malfunc_reports); result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, malfunc_reports)
}

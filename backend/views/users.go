package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AdamPekny/IIS/backend/serializers"
	utils "github.com/AdamPekny/IIS/backend/utils"
)

func Signup(ctx *gin.Context) {
	var user serializers.UserPublicSerializer

	// Validate User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if !user.Valid() {
		ctx.IndentedJSON(http.StatusBadRequest, user.ValidatorErrs)
		return
	}

	user_model := user.Create_model()

	// Create User
	result := utils.DB.Create(user_model)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

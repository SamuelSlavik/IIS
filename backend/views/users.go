package views

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AdamPekny/IIS/backend/serializers"
	utils "github.com/AdamPekny/IIS/backend/utils"
)

func Create_user(ctx *gin.Context) {
	var user serializers.UserPublicSerializer

	// Connect to db
	db, err := utils.Conn()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	// Validate User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	fmt.Print(user.Email)
	fmt.Print(user.Password)

	if !user.Valid() {
		fmt.Print("\n")
		fmt.Print(user.ValidatorErrs)
		fmt.Print("\n")
		ctx.IndentedJSON(http.StatusBadRequest, user.ValidatorErrs)
		return
	}

	user_model := user.Create_model()

	// Create User
	result := db.Create(user_model)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

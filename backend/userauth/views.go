package userauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	utils "github.com/AdamPekny/IIS/backend/utils"
)

func Create_user_type(ctx *gin.Context) {
	var user_type UserType

	if err := ctx.BindJSON(&user_type); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	db, err := utils.Conn()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	result := db.Create(&user_type)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func Create_user(ctx *gin.Context) {
	var user User

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

	// Create User
	result := db.Create(&user)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

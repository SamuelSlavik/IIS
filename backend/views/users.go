package views

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/serializers"
	utils "github.com/AdamPekny/IIS/backend/utils"
)

func Signup(ctx *gin.Context) {
	var user_serializer serializers.UserSignupSerializer

	// Validate User
	if err := ctx.BindJSON(&user_serializer); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if !user_serializer.Valid() {
		ctx.IndentedJSON(http.StatusBadRequest, user_serializer.ValidatorErrs)
		return
	}

	user_model := user_serializer.Create_model()

	// Create User
	result := utils.DB.Create(user_model)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func Login(ctx *gin.Context) {
	var user_serializer serializers.UserLoginSerializer
	var user_model models.User

	// Validate and bind User to serializer
	if err := ctx.BindJSON(&user_serializer); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if !user_serializer.Valid() {
		ctx.IndentedJSON(http.StatusBadRequest, user_serializer.ValidatorErrs)
		return
	}

	// Find user in db
	result := utils.DB.First(&user_model, "email = ?", user_serializer.Email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"UserNotFoundErr": "User not found!"})
		return
	}

	// Check password
	err := bcrypt.CompareHashAndPassword([]byte(user_model.Password), []byte(user_serializer.Password))

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"UserAuthErr": "Password does not match!"})
		return
	}

	// Create jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user_model.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	token_string, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		fmt.Print(err.Error())
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"UserAuthInternalErr": "Could not create token!"})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token_string, 3600, "", "", false, true)

	ctx.IndentedJSON(http.StatusOK, gin.H{})
}

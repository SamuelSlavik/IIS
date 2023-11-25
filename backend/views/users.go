package views

import (
	"errors"
	"net/http"
	"os"
	"strconv"
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

	user_model := user_serializer.ToModel()

	// Create User
	result := utils.DB.Create(user_model)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func DeleteUser(ctx *gin.Context) {
	user_id := ctx.Param("id")

	if user_id == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Fetch user from the database
	var user models.User
	result := utils.DB.First(&user, "id = ?", user_id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if result.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Delete user from the database
	delete_result := utils.DB.Delete(&user)
	if delete_result.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func Login(ctx *gin.Context) {
	var user_serializer serializers.UserLoginSerializer
	var user_model models.User
	var user_public serializers.UserPublicSerializer

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
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	// Check password
	err := bcrypt.CompareHashAndPassword([]byte(user_model.Password), []byte(user_serializer.Password))

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Password does not match!"})
		return
	}

	// Create jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user_model.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	token_string, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token_string, 3600, "", "", false, true)

	if err := user_public.FromModel(user_model); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, user_public)
}

func Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func ListUsers(ctx *gin.Context) {
	var user_models []models.User
	var user_serializers []serializers.UserPublicSerializer

	res := utils.DB.Find(&user_models)

	if res.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": res.Error.Error(),
		})
		return
	}

	for _, user := range user_models {
		user_serialized := serializers.UserPublicSerializer{}
		if err := user_serialized.FromModel(user); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user_serializers = append(user_serializers, user_serialized)
	}

	ctx.IndentedJSON(http.StatusOK, user_serializers)
}

func RetrieveUser(ctx *gin.Context) {
	var user_model models.User
	var user_serializer serializers.UserPublicSerializer

	uid, ok := ctx.Params.Get("id")

	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": "No UID provided!",
		})
		return
	}

	res := utils.DB.First(&user_model, uid)

	if res.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": res.Error.Error(),
		})
		return
	}

	user_serializer = serializers.UserPublicSerializer{}
	if err := user_serializer.FromModel(user_model); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, user_serializer)
}

func RetrieveCurrentUser(ctx *gin.Context) {
	user_serializer := serializers.UserPublicSerializer{}

	user_ctx, exists := ctx.Get("user")

	if !exists {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"Error": "User not found!",
		})
		return
	}

	user_model, ok := user_ctx.(models.User)

	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"Error": "Not a valid user!",
		})
		return
	}

	if err := user_serializer.FromModel(user_model); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, user_serializer)
}

func UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil || id < 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID provided",
		})
		return
	}

	uid := uint(id)

	logged_user_ctx, exists := ctx.Get("user")

	if !exists {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"Error": "User not logged in",
		})
		return
	}

	logged_user_model, ok := logged_user_ctx.(models.User)

	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"Error": "Not a valid user logged in",
		})
		return
	}

	if (logged_user_model.ID != uid && logged_user_model.Role != models.AdminRole) {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"Error": "Permission denied",
		})
		return
	}

	var user_model models.User

	result := utils.DB.First(&user_model, uid)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}

	var user_update_serializer serializers.UserUpdateSerializer

	// Validate and bind User to serializer
	if err := ctx.BindJSON(&user_update_serializer); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}


	new_user := user_update_serializer.ToModel()

	// Update User
	result = utils.DB.Model(&user_model).Updates(new_user)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	var user_public serializers.UserPublicSerializer

	if err := user_public.FromModel(user_model); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, user_public)
}

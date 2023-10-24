package middleware

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func RequireAuth(permitted_roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get cookie
		token_string, err := ctx.Cookie("Authorization")

		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, _ := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check expiration
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			}

			// Find user who token belongs to
			var user models.User
			result := utils.DB.First(&user, claims["sub"])
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// Check role
			if len(permitted_roles) > 0 {
				role_ok := false
				for _, value := range permitted_roles {
					if value == string(user.Role) {
						role_ok = true
						break
					}
				}
				if !role_ok {
					ctx.AbortWithStatus(http.StatusUnauthorized)
				}
			}

			// Attach user to request
			ctx.Set("user", user)

			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

package serializers

import "time"

type UserPublicSerializer struct {
	FirstName      string    `binding:"required"`
	LastName       string    `binding:"required"`
	Email          string    `binding:"required"`
	BirthDate      time.Time `binding:"required"`
	Password       string    `binding:"required"`
	PasswordRpt    string    `binding:"required"`
	UserTypeCdName string
}

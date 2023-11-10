package serializers

import (
	"time"

	models "github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/AdamPekny/IIS/backend/validators"
	"golang.org/x/crypto/bcrypt"
)

// User Signup Serializer

type UserSignupSerializer struct {
	FirstName     string           `binding:"required"`
	LastName      string           `binding:"required"`
	Email         string           `binding:"required"`
	BirthDate     utils.CustomDate `binding:"required"`
	Password      string           `binding:"required"`
	PasswordRpt   string           `binding:"required"`
	Role          models.Role      `binding:"required"`
	ValidatorErrs []validators.ValidatorErr
}

func (u *UserSignupSerializer) Valid() bool {
	validators.Email_validator(u.Email, &u.ValidatorErrs)
	validators.Password_match(u.Password, u.PasswordRpt, &u.ValidatorErrs)

	return len(u.ValidatorErrs) == 0
}

func (u UserSignupSerializer) copy_data(user *models.User) {
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Email = u.Email
	user.BirthDate = u.BirthDate.Time
	user.Role = u.Role
	user.Password = u.Password
}

func (u UserSignupSerializer) ToModel() *models.User {
	user := &models.User{}

	pwd_hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return nil
	}

	u.Password = string(pwd_hash)

	u.copy_data(user)

	return user
}

// User Login Serializer

type UserLoginSerializer struct {
	Email         string `binding:"required"`
	Password      string `binding:"required"`
	ValidatorErrs []validators.ValidatorErr
}

func (u *UserLoginSerializer) Valid() bool {
	validators.Email_validator(u.Email, &u.ValidatorErrs)

	return len(u.ValidatorErrs) == 0
}


// User public serializer

type UserPublicSerializer struct {
	ID 			  	uint
	FirstName     	string
	LastName 	  	string
	Email         	string `binding:"required"`
	BirthDate 		time.Time
	Role 			models.Role `binding:"required"`
	CreatedAt 		time.Time
	ValidatorErrs 	[]validators.ValidatorErr 
}

func (u *UserPublicSerializer) Valid() bool {
	validators.Email_validator(u.Email, &u.ValidatorErrs)

	return len(u.ValidatorErrs) == 0
}

func (u *UserPublicSerializer) FromModel(user models.User) {
	u.ID = user.ID
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.BirthDate = user.BirthDate
	u.Role = user.Role
	u.CreatedAt = user.CreatedAt
}

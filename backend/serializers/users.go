package serializers

import (
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

func (pub_user *UserSignupSerializer) Valid() bool {
	validators.Email_validator(pub_user.Email, &pub_user.ValidatorErrs)
	validators.Password_match(pub_user.Password, pub_user.PasswordRpt, &pub_user.ValidatorErrs)

	return len(pub_user.ValidatorErrs) == 0
}

func (pub_user UserSignupSerializer) copy_data(user *models.User) {
	user.FirstName = pub_user.FirstName
	user.LastName = pub_user.LastName
	user.Email = pub_user.Email
	user.BirthDate = pub_user.BirthDate.Time
	user.Role = pub_user.Role
	user.Password = pub_user.Password
}

func (pub_user UserSignupSerializer) Create_model() *models.User {
	user := &models.User{}

	pwd_hash, err := bcrypt.GenerateFromPassword([]byte(pub_user.Password), 14)
	if err != nil {
		return nil
	}

	pub_user.Password = string(pwd_hash)

	pub_user.copy_data(user)

	return user
}

// User Login Serializer

type UserLoginSerializer struct {
	Email         string `binding:"required"`
	Password      string `binding:"required"`
	ValidatorErrs []validators.ValidatorErr
}

func (pub_user *UserLoginSerializer) Valid() bool {
	validators.Email_validator(pub_user.Email, &pub_user.ValidatorErrs)

	return len(pub_user.ValidatorErrs) == 0
}

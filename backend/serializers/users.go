package serializers

import (
	"fmt"
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
	validators.EmailValidator(u.Email, &u.ValidatorErrs)
	validators.PasswordMatch(u.Password, u.PasswordRpt, &u.ValidatorErrs)
	validators.RoleValidator(string(u.Role), &u.ValidatorErrs)

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
	validators.EmailValidator(u.Email, &u.ValidatorErrs)

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
	validators.EmailValidator(u.Email, &u.ValidatorErrs)
	validators.RoleValidator(string(u.Role), &u.ValidatorErrs)

	return len(u.ValidatorErrs) == 0
}

func (u *UserPublicSerializer) FromModel(user models.User) (err error) {
	u.ID = user.ID
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.BirthDate = user.BirthDate
	u.Role = user.Role
	u.CreatedAt = user.CreatedAt

	return nil
}

// User update serializer

type UserUpdateSerializer struct {
	FirstName     string           
	LastName      string           
	Email         string           
	BirthDate     utils.CustomDate 
	ValidatorErrs []validators.ValidatorErr
}

func (u *UserUpdateSerializer) Valid() bool {
	if u.Email != "" {
		validators.EmailValidator(u.Email, &u.ValidatorErrs)
	}

	return len(u.ValidatorErrs) == 0
}

func (u UserUpdateSerializer) copy_data(user *models.User) {
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Email = u.Email
	user.BirthDate = u.BirthDate.Time
}

func (u UserUpdateSerializer) ToModel() *models.User {
	user := &models.User{}

	u.copy_data(user)

	return user
}

type UserMaintenanceSerializer struct {
	ID uint
	FirstName string
	LastName string
	Email string
	Role models.Role
	ValidatorErrs []validators.ValidatorErr
}

func (u *UserMaintenanceSerializer) Valid() bool {
	validators.EmailValidator(u.Email, &u.ValidatorErrs)
	validators.RoleValidator(string(u.Role), &u.ValidatorErrs)

	return len(u.ValidatorErrs) == 0
}

func (u *UserMaintenanceSerializer) FromModel(user_model *models.User) (err error) {
	u.ID = user_model.ID
	u.FirstName = user_model.FirstName
	u.LastName = user_model.LastName
	u.Email = user_model.Email
	u.Role = user_model.Role

	if ok := u.Valid(); !ok {
		return fmt.Errorf("email %s is not a valid email", u.Email)
	}
	
	return nil
}

package models

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Role string

const (
	AdminRole      Role = "admin"
	SuperuserRole  Role = "superuser"
	TechnicianRole Role = "technician"
	DispatcherRole Role = "dispatcher"
	DriverRole     Role = "driver"
)

type User struct {
	gorm.Model
	FirstName   string       `gorm:"not null"`
	LastName    string       `gorm:"not null"`
	Email       string       `gorm:"not null"` // Unique among non deleted users
	BirthDate   time.Time    `gorm:"not null"`
	Password    string       `gorm:"not null"`
	Role        Role         `gorm:"not null"`
	FullName string
	Connections []Connection `gorm:"foreignKey:DriverID"`
	MalfuncReports []MalfunctionReport `gorm:"foreignKey:CreatedByRef"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	new_values := tx.Statement.Dest.(*User)
	new_values.FullName = new_values.FirstName + " " + new_values.LastName
	
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return uniqueEmailCheck(tx, u.Email)
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// Check if the email field is being updated
	if tx.Statement.Changed("Email") {
		new_values := tx.Statement.Dest.(*User)

		return uniqueEmailCheck(tx, new_values.Email)
	}

	return nil
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	if result := tx.Model(&Connection{}).Where("driver_id = ?", u.ID).Update("driver_id", nil); result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserFromCtx(ctx *gin.Context) (*User, error) {
	if user_ctx, ok := ctx.Get("user"); ok {
		user, ok := user_ctx.(User)
	
		if !ok {
			return nil, fmt.Errorf("not a valid user")
		}
	
		return &user, nil
	} else {
		return nil, fmt.Errorf("user not in context")
	}
}

func uniqueEmailCheck(tx *gorm.DB, email string) (err error) {
	var existing_user User
	result := tx.Where("email = ?", email).Find(&existing_user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		// User with the same email already exists, return an error
		return fmt.Errorf("User with email %s already exists", email)
	}

	return nil
}

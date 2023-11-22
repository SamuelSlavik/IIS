package models

import (
	"fmt"
	"time"

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
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"not null"` // Unique among non deleted users
	BirthDate time.Time `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Role      Role      `gorm:"not null"`
}

func uniqueEmailCheck(tx *gorm.DB, email string) (err error) {
	var existing_user User
	result := tx.Where("email = ?", email).First(&existing_user)

    if result.Error == nil {
        // User with the same email already exists, return an error
        return fmt.Errorf("User with email %s already exists", email)
    } else if result.Error != gorm.ErrRecordNotFound {
        // Some other error occurred, return the error
        return result.Error
    }

	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return uniqueEmailCheck(tx, u.Email)
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	var existing_user User

	result := tx.Where("id = ?", u.ID).First(&existing_user)

    if result.Error != nil {
        // User with the same email already exists, return an error
        return fmt.Errorf("User not found")
    }

	// Check if the email field is being updated
	if tx.Statement.Changed("Email") {
		new_values := tx.Statement.Dest.(*User)

		return uniqueEmailCheck(tx, new_values.Email)
	}

	return nil
}

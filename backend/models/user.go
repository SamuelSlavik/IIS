package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement;not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	BirthDate time.Time `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Role      Role      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Role string

const (
	Admin      Role = "admin"
	Superuser  Role = "superuser"
	Technician Role = "technician"
	Dispatcher Role = "dispatcher"
	Driver     Role = "driver"
)

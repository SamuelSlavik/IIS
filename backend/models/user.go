package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey;autoIncrement;not null"`
	FirstName  string `gorm:"not null"`
	LastName   string `gorm:"not null"`
	Email      string `gorm:"not null;unique"`
	Age        uint8  `gorm:"not null"`
	Password   string `gorm:"not null"`
	Salt       string `gorm:"not null"`
	UserTypeID string
	UserType   UserType  `gorm:"references:CodeName"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

type UserType struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	CodeName string `gorm:"not null;unique"`
}

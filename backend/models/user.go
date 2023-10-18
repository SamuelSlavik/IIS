package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey;autoIncrement;not null"`
	FirstName      string    `gorm:"not null"`
	LastName       string    `gorm:"not null"`
	Email          string    `gorm:"not null;unique"`
	BirthDate      time.Time `gorm:"not null"`
	Password       string    `gorm:"not null"`
	Salt           string    `gorm:"not null"`
	UserTypeCdName string
	UserType       UserType  `gorm:"foreignKey:UserTypeCdName;references:CodeName"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

type UserType struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	CodeName string `gorm:"not null;unique"`
}

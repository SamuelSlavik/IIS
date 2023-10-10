package models

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement;not null"`
	Capacity  uint `gorm:"not null"`
	Brand     string
	ImageData []byte
}

type VehicleType struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"not null"`
}

package models

import (
	"gorm.io/gorm"
)

type Stop struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;not null"`
	Name string `gorm:"not null"`
}

type TimeBetween struct {
	gorm.Model
	// TODO: pkacka
	Time uint `gorm:"not null"`
}

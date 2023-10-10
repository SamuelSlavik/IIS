package models

import (
	"gorm.io/gorm"
)

type Line struct {
	gorm.Model
	Name        string `gorm:"primaryKey;not null"`
	InitialStop string `gorm:"not null"`
	FinalStop   string `gorm:"not null"`
}

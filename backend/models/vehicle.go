package models

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement;not null"`
}

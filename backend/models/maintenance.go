package models

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	PendingStatus Status = "pending"
	InProgressStatus Status = "in progress"
	DoneStatus Status = "done"
)

type MalfunctionReport struct {
	gorm.Model
	Description string `gorm:"not null"`
	Status Status `gorm:"not null;default:pending"`
	CreatedByRef uint `gorm:"not null"`
	CreatedBy User `gorm:"foreignKey:CreatedBy"`
	VehicleRef uint `gorm:"not null"`
	Vehicle Vehicle `gorm:"foreignKey:Vehicle"`
}

type MaintenanceRequest struct {
	gorm.Model
	Description string `gorm:"not null"`
	Status Status `gorm:"not null;default:pending"`
	Deadline time.Time
}

type MaintenanceReport struct {
	Description string `gorm:"not null"`
	Cost uint `gorm:"default:0"`
}


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
	description string `gorm:"not null"`
	status Status `gorm:"not null;default:pending"`
}

type MaintenanceRequest struct {
	gorm.Model
	description string `gorm:"not null"`
	status Status `gorm:"not null;default:pending"`
	deadline time.Time
}

type MaintenanceReport struct {
	description string `gorm:"not null"`
	cost uint `gorm:"default:0"`
}


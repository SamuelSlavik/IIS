package models

import (
	"time"
)

type Status string

const (
	PendingStatus Status = "pending"
	InProgressStatus Status = "in progress"
	DoneStatus Status = "done"
)

type MalfunctionReport struct {
	ID uint `gorm:"primaryKey;autoIncrement;not null"`
	Title string `gorm:"not null;size:100"`
	Description string `gorm:"not null"`
	CreatedByRef *uint `gorm:"not null"`
	CreatedBy *User `gorm:"foreignKey:CreatedByRef"`
	VehicleRef *string `gorm:"not null"`
	Vehicle *Vehicle `gorm:"foreignKey:VehicleRef"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	MaintenReqs []MaintenanceRequest `gorm:"foreignkey:MalfuncRepRef;constraint:OnDelete:CASCADE"`
}

type MaintenanceRequest struct {
	ID uint `gorm:"primaryKey;autoIncrement;not null"`
	Status Status `gorm:"not null;default:pending"`
	Deadline time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	MalfuncRepRef *uint `gorm:"not null"`
	MalfuncRep *MalfunctionReport `gorm:"foreignkey:MalfuncRepRef"`
	CreatedByRef *uint `gorm:"not null"`
	CreatedBy *User `gorm:"foreignKey:CreatedByRef"`
	ResolvedByRef *uint 
	ResolvedBy *User `gorm:"foreignKey:ResolvedByRef"`
	MaintenRep *MaintenanceReport `gorm:"foreignkey:MaintenReqRef;constraint:OnDelete:CASCADE"`
}

type MaintenanceReport struct {
	ID uint `gorm:"primaryKey;autoIncrement;not null"`
	Title string `gorm:"not null;size:100"`
	Description string `gorm:"not null"`
	Cost uint `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	MaintenReqRef *uint `gorm:"not null;unique"`
	MaintenReq *MaintenanceRequest `gorm:"foreignKey:MaintenReqRef"`
}


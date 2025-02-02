// package models contains gorm model definitions for ORM usage
// this file contains models for maintenance
package models

import (
	"time"
)

type Status string

const (
	PendingStatus    Status = "pending"
	InProgressStatus Status = "progress"
	DoneStatus       Status = "done"
)

type MalfunctionReport struct {
	ID           uint                 `gorm:"primaryKey;autoIncrement;not null"`
	Title        string               `gorm:"not null;size:100"`
	Description  string               `gorm:"not null"`
	CreatedByRef *uint                `gorm:"not null"`
	CreatedBy    *User                `gorm:"foreignKey:CreatedByRef"`
	VehicleRef   *string              `gorm:"not null"`
	Vehicle      *Vehicle             `gorm:"foreignKey:VehicleRef"`
	CreatedAt    time.Time            `gorm:"autoCreateTime"`
	MaintenReqs  []MaintenanceRequest `gorm:"foreignkey:MalfuncRepRef;constraint:OnDelete:CASCADE"`
}

type MaintenanceRequest struct {
	ID            uint   `gorm:"primaryKey;autoIncrement;not null"`
	Status        Status `gorm:"not null;default:pending"`
	Deadline      *time.Time
	CreatedAt     time.Time          `gorm:"autoCreateTime"`
	MalfuncRepRef *uint              `gorm:"not null"`
	MalfuncRep    *MalfunctionReport `gorm:"foreignkey:MalfuncRepRef"`
	CreatedByRef  *uint              `gorm:"not null"`
	CreatedBy     *User              `gorm:"foreignKey:CreatedByRef"`
	ResolvedByRef *uint
	ResolvedBy    *User              `gorm:"foreignKey:ResolvedByRef"`
	MaintenRep    *MaintenanceReport `gorm:"foreignkey:MaintenReqRef;constraint:OnDelete:CASCADE"`
}

type MaintenanceReport struct {
	ID            uint                `gorm:"primaryKey;autoIncrement;not null"`
	Title         string              `gorm:"not null;size:100"`
	Description   string              `gorm:"not null"`
	Cost          float64             `gorm:"default:0.0;type:decimal(11,2);not null"`
	CreatedAt     time.Time           `gorm:"autoCreateTime"`
	MaintenReqRef *uint               `gorm:"not null;unique"`
	MaintenReq    *MaintenanceRequest `gorm:"foreignKey:MaintenReqRef"`
}

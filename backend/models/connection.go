package models

import (
	"time"

	"gorm.io/gorm"
)

type Connection struct {
	gorm.Model
	ID            uint `gorm:"primaryKey;autoIncrement;not null"`
	ArrivalTime   time.Time
	DepartureTime time.Time
	// Dirrection idk aky typ lmao
}

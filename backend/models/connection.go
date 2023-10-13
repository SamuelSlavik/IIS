package models

import (
	"time"
)

type Connection struct {
	ID            uint `gorm:"primaryKey;autoIncrement;not null"`
	ArrivalTime   time.Time
	DepartureTime time.Time
	// Dirrection idk aky typ lmao
	LineRefer string
	Line      Line `gorm:"foreignKey:LineRefer"`
}

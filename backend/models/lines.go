package models

import (
	"time"
)

type Line struct {
	Name        string  `gorm:"primaryKey;unique;not null"`
	InitialStop string  `gorm:"not null"`
	FinalStop   string  `gorm:"not null"`
	Stops       []*Stop `gorm:"many2many:line_stops;"`
	Vehicles    []Vehicle
}

type Stop struct {
	ID    uint    `gorm:"primaryKey;not null;autoIncrement"`
	Name  string  `gorm:"not null"`
	Lines []*Line `gorm:"many2many:line_stops;"`
	Stop1 []*Stop `gorm:"many2many:time_betweens;"`
}

type TimeBetween struct {
	Time uint `gorm:"not null"`
}

type Connection struct {
	ID            uint `gorm:"primaryKey;autoIncrement;not null"`
	ArrivalTime   time.Time
	DepartureTime time.Time
	// Dirrection idk aky typ lmao
	LineRefer string
	Line      Line `gorm:"foreignKey:LineRefer"`
}

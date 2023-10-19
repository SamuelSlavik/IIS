package models

import (
	"time"
)

type Line struct {
	Name        string `gorm:"primaryKey;unique;not null"`
	InitialStop string `gorm:"not null"`
	FinalStop   string `gorm:"not null"`
	Connections []Connection
	Segments    []*Segment `gorm:"many2many:line_segments;"`
}

type Segment struct {
	StopName1 string  `gorm:"primaryKey;autoIncrement:false"`
	StopName2 string  `gorm:"primaryKey;autoIncrement:false"`
	Stop1     Stop    `gorm:"foreignKey:StopName1;references:Name"`
	Stop2     Stop    `gorm:"foreignKey:StopName2;references:Name"`
	Time      uint    `gorm:"not null"`
	Lines     []*Line `gorm:"many2many:line_segments;"`
}

type Stop struct {
	ID   uint   `gorm:"primaryKey;not null;autoIncrement"`
	Name string `gorm:"not null;unique"`
}

type Connection struct {
	ID            uint `gorm:"primaryKey;autoIncrement;not null"`
	ArrivalTime   time.Time
	DepartureTime time.Time
	VehicleID     uint
	// Dirrection idk aky typ lmao TODO: typ
	LineName string
}

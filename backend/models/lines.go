package models

import (
	"time"
)

type Line struct {
	Name        string `gorm:"primaryKey;unique;not null"`
	InitialStop string `gorm:"not null"`
	FinalStop   string `gorm:"not null"`
	Connections []Connection
	Segments    []Segment
}

type Segment struct {
	StopName1 string `gorm:"primaryKey;autoIncrement:false"`
	StopName2 string `gorm:"primaryKey;autoIncrement:false"`
	Stop1     Stop   `gorm:"foreignKey:StopName1;references:Name"`
	Stop2     Stop   `gorm:"foreignKey:StopName2;references:Name"`
	Time      uint   `gorm:"not null"`
	LineName  string
}

type Connection struct {
	ID                  uint `gorm:"primaryKey;autoIncrement;not null"`
	DepartureTime       time.Time
	ArrivalTime         time.Time
	Dirrection          bool   //TRUE: Initial->Final FALSE: Final->Initial
	VehicleRegistration string `gorm:"default:null"`
	LineName            string `gorm:"not null"`
	DriverID            uint   `gorm:"default:null"`
}

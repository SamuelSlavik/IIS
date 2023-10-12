package models

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	ID               uint `gorm:"primaryKey;autoIncrement;not null"`
	Capacity         uint `gorm:"not null"`
	Brand            string
	ImageData        []byte
	VehicleTypeRefer uint        //id cudzieho kluca i guess ??
	VehicleType      VehicleType `gorm:"foreignKey:VehicleTypeRefer"`
	/*LineRefer        string //TODO: IDK AKO LIKNUT STRING PKCKO
	Line             Line `gorm:"constraint:unique;foreignKey:LineRefer"`*/
}

type VehicleType struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"not null"`
}

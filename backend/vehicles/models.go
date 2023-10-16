package vehicles

import "github.com/AdamPekny/IIS/backend/lines"

type Vehicle struct {
	ID               uint `gorm:"primaryKey;autoIncrement;not null"`
	Capacity         uint `gorm:"not null"`
	Brand            string
	ImageData        []byte
	VehicleTypeRefer uint        //id cudzieho kluca i guess ??
	VehicleType      VehicleType `gorm:"foreignKey:VehicleTypeRefer"`
	LineRefer        string
	Line             lines.Line `gorm:"foreignKey:LineRefer"`
}

type VehicleType struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"not null"`
}

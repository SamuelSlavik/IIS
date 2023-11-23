package models

type Vehicle struct {
	ID              uint   `gorm:"primaryKey;autoIncrement;not null"`
	Registration    string `gorm:"unique;not null"`
	Capacity        uint   `gorm:"not null"`
	Brand           string
	ImageData       []byte
	VehicleTypeName string
	VehicleType     VehicleType `gorm:"foreignKey:VehicleTypeName;references:Type"`
	LineName        *string
	Connections     []Connection
	Malfunctions     []MalfunctionReport `gorm:"foreignKey:Vehicle"`
}

type VehicleType struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;not null"`
	Type string `gorm:"unique;not null"`
}

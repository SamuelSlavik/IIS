package models

type Vehicle struct {
	ID              uint `gorm:"primaryKey;autoIncrement;not null"`
	Capacity        uint `gorm:"not null"`
	Brand           string
	ImageData       []byte
	VehicleTypeName string      //id cudzieho kluca i guess ??
	VehicleType     VehicleType `gorm:"foreignKey:VehicleTypeName;references:Type"`
	Connections     []Connection
}

type VehicleType struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;not null"`
	Type string `gorm:"unique;not null"`
}

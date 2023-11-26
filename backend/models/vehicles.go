package models

type Vehicle struct {
	Registration    string `gorm:"primaryKey;unique;not null"`
	Capacity        uint   `gorm:"not null"`
	Brand           string
	ImageData       []byte
	VehicleTypeName string
	VehicleType     VehicleType         `gorm:"foreignKey:VehicleTypeName;references:Type"`
	Connections     []Connection        `gorm:"constraint:OnDelete:SET NULL"`
	Malfunctions    []MalfunctionReport `gorm:"foreignKey:VehicleRef;constraint:OnDelete:CASCADE"`
}

type VehicleType struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;not null"`
	Type string `gorm:"unique;not null"`
}

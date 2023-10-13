package models

type Line struct {
	Name        string  `gorm:"primaryKey;unique;not null"`
	InitialStop string  `gorm:"not null"`
	FinalStop   string  `gorm:"not null"`
	Stops       []*Stop `gorm:"many2many:line_stops;"`
}

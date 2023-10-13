package models

type Stop struct {
	ID    uint    `gorm:"primaryKey;not null"`
	Name  string  `gorm:"not null"`
	Lines []*Line `gorm:"many2many:line_stops;"`
}

type TimeBetween struct {
	// TODO: pkacka
	Time uint `gorm:"not null"`
}

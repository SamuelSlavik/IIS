package models

type Stop struct {
	ID    uint    `gorm:"primaryKey;not null;autoIncrement"`
	Name  string  `gorm:"not null"`
	Lines []*Line `gorm:"many2many:line_stops;"`
	Stop1 []*Stop `gorm:"many2many:time_betweens;"`
}

type TimeBetween struct {
	Time uint `gorm:"not null"`
}

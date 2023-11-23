package models

type Stop struct {
	ID   uint   `gorm:"primaryKey;not null;autoIncrement"`
	Name string `gorm:"not null;unique"`
}

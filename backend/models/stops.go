// package models contains gorm model definitions for ORM usage
// this file contains models for stops
package models

type Stop struct {
	ID   uint   `gorm:"primaryKey;not null;autoIncrement"`
	Name string `gorm:"not null;unique"`
}

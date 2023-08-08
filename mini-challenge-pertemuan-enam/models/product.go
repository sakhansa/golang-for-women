package models

import "time"

type Product struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null"`
	Variants  []Variant
	CreatedAt time.Time
	UpdatedAt time.Time
}

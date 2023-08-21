package models

import "time"

type Student struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Age       int    `gorm:"not null"`
	Scores    []Score
	CreatedAt time.Time
	UpdatedAt time.Time
}

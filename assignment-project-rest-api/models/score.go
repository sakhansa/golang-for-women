package models

import "time"

type Score struct {
	ID              uint   `gorm:"primary_key"`
	AssignmentTitle string `gorm:"not null;type:varchar(191)" json:"assignment_title"`
	Description     string `gorm:"not null;type:varchar(191)"`
	Score           int    `gorm:"not null"`
	StudentId       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

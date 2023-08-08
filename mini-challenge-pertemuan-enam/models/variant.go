package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Variant struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	VariantName string `gorm:"not null"`
	Stock       int    `gorm:"not null"`
	ProductID   uint   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Book before create()")

	if len(v.VariantName) < 5 {
		err = errors.New("Variant Name is too short. Minimum length of Variant Name is 5")
	}

	if v.Stock == 0 {
		err = errors.New("Stock is too small. Minimum stock is 1")
	}

	return
}

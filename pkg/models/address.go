package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Entity
	Direction string `gorm:"not null"`
	City      string `gorm:"not null"`
	Number    int64  `gorm:"not null"`
	Zipcode   int64  `gorm:"not null"`
}

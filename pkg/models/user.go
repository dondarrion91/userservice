package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Entity
	Fullname *string `gorm:"not null"`
	Email    string  `gorm:"unique;not null"`
}

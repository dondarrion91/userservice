package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Entity
	Name *string `gorm:"not null"`
}

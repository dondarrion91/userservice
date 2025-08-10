package models

import (
	"time"
)

type Entity struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

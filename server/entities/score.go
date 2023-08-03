package entities

import (
	"time"
)

type Score struct {
	ID        int       `gorm:"autoIncrement;"`
	CreatedAt time.Time `gorm:"not null;"`
	Score     int       `gorm:"not null;"`
	UserID    string    `gorm:"not null;"`
}

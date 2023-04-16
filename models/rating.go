package models

import "time"

type Rating struct {
	ID            string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	DestinationID string `gorm:"size:36;not null;index"`
	Rate          int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

package models

import "time"

type ImageDestination struct {
	ID            string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	FileName      string
	IsPrimary     bool
	DestinationID string `gorm:"size:36;not null;index"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

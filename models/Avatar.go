package models

import "time"

type Avatar struct {
	ID        string `gorm:"size:36;not null;uniqueIndex"`
	UserID    string `gorm:"size:36;not null;index"`
	FileName  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

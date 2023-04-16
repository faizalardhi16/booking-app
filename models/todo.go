package models

import "time"

type Todo struct {
	ID        string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name      string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

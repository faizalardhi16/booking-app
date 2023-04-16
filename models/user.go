package models

import "time"

type User struct {
	ID        string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;not null;unique"`
	Password  string `gorm:"size:100;not null"`
	Avatar    Avatar `gorm:"foreignKey:UserID;reference:id"`
	Role      string `gorm:"size:100;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

package models

import "time"

type Destination struct {
	ID               string             `gorm:"size:36;not null;uniqueIndex;primary_key"`
	DestinationName  string             `gorm:"type:varchar(255);not null"`
	Description      string             `gorm:"type:text;not null"`
	ShortDescription string             `gorm:"type:text;not null"`
	Rating           []Rating           `gorm:"one2many:rate_dest"`
	ImageDestination []ImageDestination `gorm:"one2many:image_dest"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

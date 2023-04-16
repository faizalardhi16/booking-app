package destination

import "time"

type Destination struct {
	ID               string
	DestinationName  string
	Description      string
	ShortDescription string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Rating struct {
	ID            string
	DestinationID string
	Rate          int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ImageDestination struct {
	ID            string
	FileName      string
	IsPrimary     int
	DestinationID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

package destination

import "time"

type DestinationDto struct {
	ID               string
	DestinationName  string
	Description      string
	ShortDescription string
	Rating           float64
	ImageDestination []ImageDestinationDto
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ImageDestinationDto struct {
	FileName  string
	IsPrimary bool
}

type RatingDto struct {
	TotalRate float64
}

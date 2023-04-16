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
	FileName  string `json:"fileName"`
	IsPrimary bool   `json:"isPrimary"`
	Url       string `json:"url"`
}

type RatingDto struct {
	TotalRate float64
}

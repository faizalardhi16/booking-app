package destination

type DestinationResponse struct {
	ID               string                `json:"id"`
	DestinationName  string                `json:"destination_name"`
	Description      string                `json:"description"`
	ShortDescription string                `json:"shortDescription"`
	Image            []ImageDestinationDto `json:"imageDestination"`
	Rating           float64               `json:"rating"`
}

func FormatCreateDestination(destination Destination) DestinationResponse {
	res := DestinationResponse{}

	res.ID = destination.ID
	res.Description = destination.Description
	res.DestinationName = destination.DestinationName
	res.ShortDescription = destination.ShortDescription
	res.Image = []ImageDestinationDto{}
	res.Rating = 0.0

	return res
}

func FormatGetAllDestination(db []DestinationDto) []DestinationResponse {
	res := []DestinationResponse{}

	for _, t := range db {
		r := DestinationResponse{}

		r.ID = t.ID
		r.DestinationName = t.DestinationName
		r.Description = t.Description
		r.ShortDescription = t.ShortDescription
		r.Image = t.ImageDestination
		r.Rating = t.Rating

		res = append(res, r)
	}

	return res
}

package destination

import validation "github.com/go-ozzo/ozzo-validation"

type DestinationInput struct {
	DestinationName  string `json:"destination_name" binding:"required"`
	Description      string `json:"description" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
}

type RatingInput struct {
	Rate          int64  `json:"rate" binding:"required"`
	DestinationID string `json:"destination_id" binding:"required"`
}

type ImageInput struct {
	FileName      string `json:"file_name" binding:"required"`
	IsPrimary     bool   `json:"is_primary" binding:"required"`
	DestinationID string `json:"destination_id" binding:"required"`
}

func (p DestinationInput) ValidationDestinationInput() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.DestinationName, validation.Required),
		validation.Field(&p.Description, validation.Required),
		validation.Field(&p.ShortDescription, validation.Required),
	)
}

func (p RatingInput) ValidateRatingInput() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Rate, validation.Required),
		validation.Field(&p.DestinationID, validation.Required, validation.Length(0, 36)),
	)
}

func (p ImageInput) ValidationImageInput() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.FileName, validation.Required),
		validation.Field(&p.IsPrimary, validation.Required),
		validation.Field(&p.DestinationID, validation.Required, validation.Length(0, 36)),
	)
}

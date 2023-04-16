package destination

import validation "github.com/go-ozzo/ozzo-validation"

type DestinationInput struct {
	DestinationName  string `json:"destinationName" binding:"required"`
	Description      string `json:"description" binding:"required"`
	ShortDescription string `json:"shortDescription" binding:"required"`
}

type RatingInput struct {
	Rate          int64  `json:"rate" binding:"required"`
	DestinationID string `json:"destinationId" binding:"required"`
}

type ImageInput struct {
	FileName      string `json:"fileName" binding:"required"`
	IsPrimary     bool   `json:"isPrimary" binding:"required"`
	DestinationID string `json:"destinationId" binding:"required"`
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

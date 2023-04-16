package destination

import "github.com/google/uuid"

type Service interface {
	FindAllDestination() ([]DestinationDto, error)
	SaveDestination(input DestinationInput) (Destination, error)
	SaveRating(input RatingInput) (string, error)
	SaveImage(input ImageInput) (string, error)
}

type service struct {
	repository Repository
}

func NewDestinationService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAllDestination() ([]DestinationDto, error) {
	data, err := s.repository.GetAllDestination()

	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *service) SaveDestination(input DestinationInput) (Destination, error) {
	req := Destination{}

	req.ID = uuid.New().String()
	req.DestinationName = input.DestinationName
	req.Description = input.Description
	req.ShortDescription = input.ShortDescription

	createData, err := s.repository.CreateDestination(req)

	if err != nil {
		return createData, err
	}

	return createData, nil

}

func (s *service) SaveRating(input RatingInput) (string, error) {
	req := Rating{}

	req.ID = uuid.New().String()
	req.Rate = int(input.Rate)
	req.DestinationID = input.DestinationID

	rate, err := s.repository.MakeRating(req)

	if err != nil {
		return rate, err
	}

	return rate, nil
}

func (s *service) SaveImage(image ImageInput) (string, error) {
	req := ImageDestination{}
	isPrimary := 0

	if image.IsPrimary != false {
		isPrimary = 1
	}

	req.ID = uuid.New().String()
	req.DestinationID = image.DestinationID
	req.IsPrimary = isPrimary
	req.FileName = image.FileName

	upload, err := s.repository.UploadImage(req)

	if err != nil {
		return upload, err
	}

	return upload, nil
}

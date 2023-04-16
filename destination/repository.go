package destination

import (
	"bookingApp/function"

	"gorm.io/gorm"
)

type Repository interface {
	CreateDestination(destination Destination) (Destination, error)
	GetAllDestination() ([]DestinationDto, error)
	MakeRating(rating Rating) (string, error)
	UploadImage(image ImageDestination) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewDestinationRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateDestination(destination Destination) (Destination, error) {
	err := r.db.Create(&destination).Error

	if err != nil {
		return destination, err
	}

	return destination, nil
}

func (r *repository) GetAllDestination() ([]DestinationDto, error) {
	var getDestination []Destination
	var destination []DestinationDto
	var count float64

	err := r.db.Raw("select * from destinations").Scan(&getDestination).Error

	if err != nil {
		return destination, err
	}

	for _, t := range getDestination {
		var image []ImageDestinationDto
		res := DestinationDto{}

		res.ID = t.ID
		res.Description = t.Description
		res.DestinationName = t.DestinationName
		res.ShortDescription = t.ShortDescription
		res.CreatedAt = t.CreatedAt
		res.UpdatedAt = t.UpdatedAt

		row := r.db.Table("ratings").Where("destination_id = ?", t.ID).Select("avg(rate)").Row()
		row.Scan(&count)

		res.Rating = count

		err := r.db.Raw("select file_name, is_primary from image_destinations where destination_id = ?", t.ID).
			Scan(&image).Error

		if err != nil {
			return destination, err
		}

		index := 0

		for _, r := range image {
			url, err := function.GetPresignedUrl(r.FileName)

			if err != nil {
				return destination, err
			}

			image[index].Url = url

			index++
		}

		res.ImageDestination = image

		destination = append(destination, res)
	}

	return destination, nil

}

func (r *repository) MakeRating(rating Rating) (string, error) {
	err := r.db.Create(&rating).Error

	if err != nil {
		return "Failed", err
	}

	return "Success", nil

}

func (r *repository) UploadImage(image ImageDestination) (string, error) {
	err := r.db.Create(&image).Error

	if err != nil {
		return "Failed", err
	}

	return "Success", nil
}

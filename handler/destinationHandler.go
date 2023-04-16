package handler

import (
	"bookingApp/constant"
	"bookingApp/destination"
	"bookingApp/helper"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
)

type destinationHandler struct {
	destinationService destination.Service
	svc                *s3.S3
}

func NewDestinationHandler(destinationService destination.Service, svc *s3.S3) *destinationHandler {
	return &destinationHandler{destinationService, svc}
}

func (h *destinationHandler) GetAllDestination(c *fiber.Ctx) error {

	findAll, err := h.destinationService.FindAllDestination()

	if err != nil {
		response := helper.APIResponse(failed, http.StatusBadRequest, nil, "Failed to load data")
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	format := destination.FormatGetAllDestination(findAll)

	response := helper.APIResponse(success, http.StatusOK, format, "Success to load the data")

	return c.Status(http.StatusOK).JSON(response)
}

func (h *destinationHandler) CreateDestination(c *fiber.Ctx) error {
	var input destination.DestinationInput

	err := c.BodyParser(&input)

	if err != nil {
		response := helper.APIResponse(failed, http.StatusMethodNotAllowed, nil, "Failed to save data, method not allowed")
		return c.Status(http.StatusMethodNotAllowed).JSON(response)
	}

	newDestination, err := h.destinationService.SaveDestination(input)

	if err != nil {
		response := helper.APIResponse(failed, http.StatusBadRequest, nil, "Failed to save data")
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	format := destination.FormatCreateDestination(newDestination)

	response := helper.APIResponse(success, http.StatusOK, format, "Success to Save Data")

	return c.Status(http.StatusOK).JSON(response)
}

func (h *destinationHandler) CreateRating(c *fiber.Ctx) error {
	var input destination.RatingInput

	err := c.BodyParser(&input)

	if err != nil {
		response := helper.APIResponse(failed, http.StatusMethodNotAllowed, constant.ConditionEnum.Failed, "Failed to make rating")
		return c.Status(http.StatusMethodNotAllowed).JSON(response)
	}

	result, err := h.destinationService.SaveRating(input)

	if err != nil {
		response := helper.APIResponse(failed, http.StatusBadRequest, constant.ConditionEnum.Failed, "Failed to make rating")
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse(success, http.StatusOK, result, "Success to make rating")

	return c.Status(http.StatusOK).JSON(response)

}

func (h *destinationHandler) CreateImage(c *fiber.Ctx) error {

	file, err := c.FormFile("upload")

	if err != nil {
		response := helper.APIResponse(failed, http.StatusBadRequest, constant.ConditionEnum.Failed, "Failed to Upload Image")
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	isPrimary := c.FormValue("isPrimary")
	booleanValue, err := strconv.ParseBool(isPrimary)
	if err != nil {
		response := helper.APIResponse(failed, http.StatusBadRequest, constant.ConditionEnum.Failed, "Failed to Upload Image")
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	f, err := file.Open()

	if err != nil {
		return err
	}

	//upload to s3
	_, err = h.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("mygobucketzores"),
		Key:    aws.String(file.Filename),
		Body:   f,
	})

	if err != nil {
		response := helper.APIResponse(failed, http.StatusBadRequest, constant.ConditionEnum.Failed, "Failed Upload Image to AWS")
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	upload, err := h.destinationService.SaveImage(destination.ImageInput{
		IsPrimary:     booleanValue,
		FileName:      file.Filename,
		DestinationID: c.FormValue("destination_id"),
	})

	if err != nil {
		response := helper.APIResponse(failed, http.StatusBadRequest, constant.ConditionEnum.Failed, "Failed Upload Image to DB")
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse(success, http.StatusOK, upload, "Success to Upload Image")

	return c.Status(http.StatusOK).JSON(response)

}

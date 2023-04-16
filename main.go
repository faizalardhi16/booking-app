package main

import (
	"bookingApp/destination"
	"bookingApp/function"
	"bookingApp/handler"
	"bookingApp/models"
	"bookingApp/todo"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		envs["DB_USER"],
		envs["DB_PASSWORD"],
		envs["DB_HOST"],
		envs["DB_NAME"],
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	models.RegisterDB(db)

	//setup s3 uploader
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(envs["AWS_REGION"]), // replace with your S3 bucket region
		Credentials: credentials.NewStaticCredentials(
			envs["AWS_ACCESS_KEY_ID"],     // replace with your AWS access key ID
			envs["AWS_SECRET_KEY_ACCESS"], // replace with your AWS secret access key
			""),
	})

	if err != nil {
		fmt.Println("Failed to create session:", err)
		os.Exit(1)
	}

	svc := s3.New(sess)

	//repository
	todoRepo := todo.NewRepository(db)
	destinationRepo := destination.NewDestinationRepository(db)

	//service
	todoService := todo.NewServiceTodo(todoRepo)
	destinationService := destination.NewDestinationService(destinationRepo)

	//handler
	todoHandler := handler.NewTodoHandler(todoService)
	destinationHandler := handler.NewDestinationHandler(destinationService, svc)

	api := app.Group("/api/v1")
	//end setup s3 uploader

	//GET method
	api.Get("/todo", todoHandler.GetAllTodoHandler)
	api.Get("/destination", destinationHandler.GetAllDestination)

	//POST method
	api.Post("/todo", todoHandler.CreateTodoHandler)
	api.Post("/rating", destinationHandler.CreateRating)
	api.Post("/destination", destinationHandler.CreateDestination)
	api.Post("/file", func(c *fiber.Ctx) error {

		// url, err := function.GetPresignedUrl(c.FormValue("fileName"))

		// if err != nil {
		// 	return err
		// }

		upload, err := c.FormFile("upload")

		if err != nil {
			return c.JSON(fiber.Map{
				"Message":     "Failed to upload file",
				"Status":      400,
				"Data":        nil,
				"Acknowledge": 1,
			})
		}

		file, err := function.UploadToStorage(upload)

		if err != nil {
			return c.JSON(fiber.Map{
				"Message":     "Failed to upload file",
				"Status":      400,
				"Data":        nil,
				"Acknowledge": 1,
			})
		}

		return c.JSON(fiber.Map{
			"Message":     "Success to upload file",
			"Status":      201,
			"Data":        file,
			"Acknowledge": 1,
		})
	})
	api.Post("/upload", destinationHandler.CreateImage)

	app.Listen(":8080")
}

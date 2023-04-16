package function

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

type BlobService interface {
	UploadToStorage(file *multipart.FileHeader) (string, error)
	GetPresignedUrl(fileName string) (string, error)
}

func UploadToStorage(file *multipart.FileHeader) (string, error) {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
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

	//open file

	f, err := file.Open()

	if err != nil {
		return "", err
	}

	//upload to s3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("mygobucketzores"),
		Key:    aws.String(file.Filename),
		Body:   f,
	})

	if err != nil {
		return "", err
	}

	return file.Filename, nil

}

func GetPresignedUrl(fileName string) (string, error) {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
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

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("mygobucketzores"),
		Key:    aws.String(fileName),
	})

	if err != nil {
		return "", err
	}

	url, err := req.Presign(15 * time.Minute)

	if err != nil {
		return "", err
	}

	return url, nil
}

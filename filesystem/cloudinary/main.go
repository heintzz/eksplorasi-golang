package main

import (
	"bytes"
	"cloudinary/pkg/images"
	"cloudinary/utils/config"
	"context"
	"io"
	"log"
	"os"
)

type CloudService interface {
	// @Param file refer to file buffer
	// @Param pathDestination refer to target directory/bucket in cloud provider
	Upload(ctx context.Context, file interface{}, pathDestination string) (uri string, err error)
	Remove(ctx context.Context, path string) (err error)
}

type Services struct {
	cloud CloudService
}

func main() {
	err := config.LoadConfig(".env")
	if err != nil {
		log.Println("error when try to LoadConfig with detail :", err.Error())
	}
	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cloudProvider := images.NewCloudinary(cloudName, apiKey, apiSecret)

	svc := Services{
		cloud: cloudProvider,
	}

	imagePath := "pkg/images/hi.png"
	file, err := os.Open(imagePath)
	if err != nil {
		log.Println("error when try to open file with detail :", err.Error())
		return
	}

	defer file.Close()

	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, file)
	if err != nil {
		log.Println("error when try to Copy file to buffer with detail :", err.Error())

		return
	}

	url, err := svc.cloud.Upload(context.Background(), buffer, "images")
	if err != nil {
		log.Println("error when try to Upload with detail :", err.Error())
	} else {
		log.Println("upload success with URL :", url)
	}
}

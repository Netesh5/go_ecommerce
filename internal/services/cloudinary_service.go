package services

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

// CloudinaryService wraps the config.Cloudinary struct
type CloudinaryService struct {
	CloudName string
	APIKey    string
	APISecret string
}

var cloudinaryService *CloudinaryService

func InitCloudinary(cloudName, apiKey, apiSecret string) {
	log.Println("CloudName: ", cloudName)
	cloudinaryService = &CloudinaryService{
		CloudName: cloudName,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

// UploadImageToCloudinary uploads an image to cloudinary and returns the secure URL
func UploadImageToCloudinary(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	log.Println("CloudName: ", cloudinaryService.CloudName)
	cld, err := cloudinary.NewFromParams(cloudinaryService.CloudName, cloudinaryService.APIKey, cloudinaryService.APISecret)
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: fileHeader.Filename,
		Folder:   "ecommerce/products",
	})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}

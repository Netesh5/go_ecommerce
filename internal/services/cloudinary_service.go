package services

import (
	"context"
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

func InitCloudinary(cloudName, apiKey, apiSecret string) *CloudinaryService {
	return &CloudinaryService{
		CloudName: cloudName,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

// UploadImageToCloudinary uploads an image to cloudinary and returns the secure URL
func (s *CloudinaryService) UploadImageToCloudinary(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {

	cld, err := cloudinary.NewFromParams(s.CloudName, s.APIKey, s.APISecret)
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

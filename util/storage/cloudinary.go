package storage

import (
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

// docs: https://cloudinary.com/documentation/go_integration

type CloudinaryUploader struct {
	cloudinary cloudinary.Cloudinary
}

func NewCloudinaryUploader(cloudName, apiKey, apiSecret string) (Uploader, error) {

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, fmt.Errorf("cloudinary initialize failed: %s", err.Error())
	}

	return &CloudinaryUploader{cloudinary: *cld}, nil
}

// TODO: update file field for according base64
func (cloudinaryUploader *CloudinaryUploader) Upload(ctx *gin.Context, file string) {

	// Upload the my_picture.jpg image and set the PublicID to "my_image".

	resp, err := cloudinaryUploader.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: "my_image"})
	_ = resp
	_ = err
}

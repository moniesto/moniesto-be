package storage

import (
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	"github.com/moniesto/moniesto-be/model"
)

// docs: https://cloudinary.com/documentation/go_integration

type CloudinaryUploader struct {
	cloudinary cloudinary.Cloudinary
}

// NewCloudinaryUploader creates new cloudinary uploader instance
func NewCloudinaryUploader(cloudUrl string) (Uploader, error) {
	cld, err := cloudinary.NewFromURL(cloudUrl)
	if err != nil {
		return nil, fmt.Errorf("cloudinary initialize failed: %s", err.Error())
	}

	return &CloudinaryUploader{cloudinary: *cld}, nil
}

// UploadProfilePhoto upload profile photo and return url & thumbnail_url of it
func (cloudinaryUploader *CloudinaryUploader) UploadProfilePhoto(ctx *gin.Context, base64 string) (model.ProfilePhoto, error) {
	publicID := core.CreateID()
	publicIDThumbnail := publicID + "_thumbnail"

	profilePhotoResult := model.ProfilePhoto{}

	// STEP: upload profile photo
	params := uploader.UploadParams{
		PublicID: publicID,
		Folder:   profilePhotoFolderName,
	}

	result, err := cloudinaryUploader.cloudinary.Upload.Upload(ctx, base64, params)
	if err != nil {
		return model.ProfilePhoto{}, fmt.Errorf("upload profile photo failed: %s", err.Error())
	}
	profilePhotoResult.URL = result.SecureURL

	// STEP: upload profile photo thumnbnail
	thumbnailParams := uploader.UploadParams{
		PublicID:       publicIDThumbnail,
		Folder:         profilePhotoThumbnailFolderName,
		Transformation: "h_" + profilePhotothumnbnailHeight,
	}

	result, err = cloudinaryUploader.cloudinary.Upload.Upload(ctx, base64, thumbnailParams)
	if err != nil {
		return model.ProfilePhoto{}, fmt.Errorf("upload profile photo thumbnail failed: %s", err.Error())
	}
	profilePhotoResult.ThumbnailURL = result.SecureURL

	return profilePhotoResult, nil
}

// UploadBackgroundPhoto upload background photo and return url & thumbnail_url of it
func (cloudinaryUploader *CloudinaryUploader) UploadBackgroundPhoto(ctx *gin.Context, base64 string) (model.BackgroundPhoto, error) {
	publicID := core.CreateID()
	publicIDThumbnail := publicID + "_thumbnail"

	backgroundPhotoResult := model.BackgroundPhoto{}

	// STEP: upload background photo
	params := uploader.UploadParams{
		PublicID: publicID,
		Folder:   backgroundPhotoFolderName,
	}

	result, err := cloudinaryUploader.cloudinary.Upload.Upload(ctx, base64, params)
	if err != nil {
		return model.BackgroundPhoto{}, fmt.Errorf("upload background photo failed: %s", err.Error())
	}
	backgroundPhotoResult.URL = result.SecureURL

	// STEP: upload background photo thumnbnail
	thumbnailParams := uploader.UploadParams{
		PublicID:       publicIDThumbnail,
		Folder:         backgroundPhotoThumbnailFolderName,
		Transformation: "h_" + backgroundPhotothumnbnailHeight,
	}

	result, err = cloudinaryUploader.cloudinary.Upload.Upload(ctx, base64, thumbnailParams)
	if err != nil {
		return model.BackgroundPhoto{}, fmt.Errorf("upload background photo thumbnail failed: %s", err.Error())
	}
	backgroundPhotoResult.ThumbnailURL = result.SecureURL

	return backgroundPhotoResult, nil
}

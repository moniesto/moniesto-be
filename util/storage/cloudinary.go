package storage

import (
	"fmt"
	"net/url"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	"github.com/moniesto/moniesto-be/model"
)

// docs: https://cloudinary.com/documentation/go_integration

func isURL(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

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

	// STEP: check the str is already a url -> yes -> return url
	if isURL(base64) {
		return model.ProfilePhoto{
			URL:          base64,
			ThumbnailURL: base64,
		}, nil
	}

	publicID := core.CreateID()
	publicIDThumbnail := publicID + "_thumbnail"

	profilePhotoResult := model.ProfilePhoto{}

	// STEP: upload profile photo
	params := uploader.UploadParams{
		PublicID:       publicID,
		Folder:         profilePhotoFolderName,
		Transformation: "h_" + profilePhotoHeight,
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

	// STEP: check the str is already a url -> yes -> return url
	if isURL(base64) {
		return model.BackgroundPhoto{
			URL:          base64,
			ThumbnailURL: base64,
		}, nil
	}

	publicID := core.CreateID()
	publicIDThumbnail := publicID + "_thumbnail"

	backgroundPhotoResult := model.BackgroundPhoto{}

	// STEP: upload background photo
	params := uploader.UploadParams{
		PublicID:       publicID,
		Folder:         backgroundPhotoFolderName,
		Transformation: "h_" + backgroundPhotoHeight,
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

// UploadPostDescriptionPhoto upload post description photo and return url
func (cloudinaryUploader *CloudinaryUploader) UploadPostDescriptionPhoto(ctx *gin.Context, base64 string) (model.PostPhoto, error) {
	publicID := core.CreateID()

	params := uploader.UploadParams{
		PublicID: publicID,
		Folder:   postDescriptionPhotoFolderName,
	}

	result, err := cloudinaryUploader.cloudinary.Upload.Upload(ctx, base64, params)
	if err != nil {
		return model.PostPhoto{}, fmt.Errorf("upload post description photo: %s", err.Error())
	}

	return model.PostPhoto{
		URL: result.SecureURL,
	}, nil
}

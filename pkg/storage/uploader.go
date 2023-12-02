package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
)

type Uploader interface {
	UploadProfilePhoto(*gin.Context, string) (model.ProfilePhoto, error)
	UploadBackgroundPhoto(*gin.Context, string) (model.BackgroundPhoto, error)
	UploadPostDescriptionPhoto(*gin.Context, string) (model.PostPhoto, error)
}

const profilePhotoFolderName string = "Photo/ProfilePhotos"
const profilePhotoThumbnailFolderName string = "Photo/ProfilePhotosThumbnail"
const backgroundPhotoFolderName string = "Photo/BackgroundPhotos"
const backgroundPhotoThumbnailFolderName string = "Photo/BackgroundPhotosThumbnail"

const postDescriptionPhotoFolderName string = "Photo/PostPhotos"

const profilePhotoHeight = "500"
const backgroundPhotoHeight = "800"

const profilePhotothumnbnailHeight = "100"
const backgroundPhotothumnbnailHeight = "100"

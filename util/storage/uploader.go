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

var profilePhotoFolderName string = "Photo/ProfilePhotos"
var profilePhotoThumbnailFolderName string = "Photo/ProfilePhotosThumbnail"
var backgroundPhotoFolderName string = "Photo/BackgroundPhotos"
var backgroundPhotoThumbnailFolderName string = "Photo/BackgroundPhotosThumbnail"

var postDescriptionPhotoFolderName string = "Photo/PostPhotos"

var profilePhotothumnbnailHeight = "100"
var backgroundPhotothumnbnailHeight = "100"

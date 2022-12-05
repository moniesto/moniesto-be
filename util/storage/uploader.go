package storage

import (
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type Uploader interface {
	Upload(ctx *gin.Context, file string) (*uploader.UploadResult, error)
}

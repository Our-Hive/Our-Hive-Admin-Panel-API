package api

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IUploadServicePort interface {
	UploadFile(imageData *model.ImageData, contentType string) (imageUrl string, err error)
}

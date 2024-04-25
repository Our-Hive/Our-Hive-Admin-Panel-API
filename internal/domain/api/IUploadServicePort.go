package api

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IUploadServicePort interface {
	UploadFile(imageData *model.ImageData) (imageUrl string, err error)
}

package api

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IImageDataServicePort interface {
	SaveImageData(imageData *model.ImageData) (string, *model.ImageData, error)
}

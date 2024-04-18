package spi

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IImagePersistencePort interface {
	SaveImageInDatabase(image *model.Image) error
}

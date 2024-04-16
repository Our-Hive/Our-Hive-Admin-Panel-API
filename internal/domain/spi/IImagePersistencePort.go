package spi

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IImagePersistencePort interface {
	SaveImage(image *model.Image) error
}

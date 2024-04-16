package spi

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IImageStoragePort interface {
	SaveImage(image *model.ImageData) error
}

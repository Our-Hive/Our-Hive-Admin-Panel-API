package spi

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IImagePersistencePort interface {
	SaveImageInDatabase(image *model.Image) error
	GetImageFromDatabaseByName(fileName string) (*model.Image, error)
	GetAllImagesFromDatabase(pageSize int, startAfter string) ([]*model.Image, error)
	GetImageFromDatabaseById(id string) (*model.Image, error)
}

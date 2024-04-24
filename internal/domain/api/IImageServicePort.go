package api

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IImageServicePort interface {
	SaveImage(image *model.Image) error
	GetImageByName(fileName string) (*model.Image, error)
	GetAllImages(pageSize int, startAfter string) ([]*model.Image, error)
}

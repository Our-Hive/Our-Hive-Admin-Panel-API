package usecase

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/spi"
)

type ImageUseCase struct {
	imagePersistencePort spi.IImagePersistencePort
}

func NewImageUseCase(imagePersistencePort spi.IImagePersistencePort) *ImageUseCase {
	return &ImageUseCase{imagePersistencePort: imagePersistencePort}
}

func (i ImageUseCase) SaveImage(image *model.Image) error {
	err := i.imagePersistencePort.SaveImageInDatabase(image)

	if err != nil {
		return err
	}

	return nil
}

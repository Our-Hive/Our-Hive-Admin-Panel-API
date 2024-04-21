package usecase

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/domainerror"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/spi"
)

// ImageUseCase implements IImageServicePort
type ImageUseCase struct {
	imagePersistencePort spi.IImagePersistencePort
}

func NewImageUseCase(imagePersistencePort spi.IImagePersistencePort) *ImageUseCase {
	return &ImageUseCase{imagePersistencePort: imagePersistencePort}
}

func (i ImageUseCase) SaveImage(image *model.Image) error {
	foundImage, _ := i.GetImageByName(image.Name)

	if foundImage != nil {
		return &domainerror.ImageAlreadyExistsError{FileName: image.Name}
	}

	err := i.imagePersistencePort.SaveImageInDatabase(image)

	if err != nil {
		return err
	}

	return nil
}

func (i ImageUseCase) GetImageByName(fileName string) (*model.Image, error) {
	image, err := i.imagePersistencePort.GetImageFromDatabaseByName(fileName)

	if err != nil {
		return nil, err
	}

	if image == nil {
		return nil, &domainerror.ImageNotFoundError{FileName: fileName}
	}

	return image, nil
}

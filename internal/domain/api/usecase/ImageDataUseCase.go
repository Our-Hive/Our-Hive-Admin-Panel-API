package usecase

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/spi"
)

type ImageDataUseCase struct {
	imageStoragePort spi.IImageStoragePort
}

func NewImageDataServiceUseCase(imageStoragePort spi.IImageStoragePort) *ImageDataUseCase {
	return &ImageDataUseCase{imageStoragePort: imageStoragePort}
}

func (i ImageDataUseCase) SaveImageData(imageData *model.ImageData) (string, *model.ImageData, error) {
	imageUrl, err := i.imageStoragePort.SaveImageInStorage(imageData)

	if err != nil {
		return "", nil, err
	}

	return imageUrl, imageData, nil
}

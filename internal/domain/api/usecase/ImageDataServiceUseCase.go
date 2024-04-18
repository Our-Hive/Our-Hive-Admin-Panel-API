package usecase

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/spi"
)

type ImageDataServiceUseCase struct {
	imageStoragePort spi.IImageStoragePort
}

func NewImageDataServiceUseCase(imageStoragePort spi.IImageStoragePort) *ImageDataServiceUseCase {
	return &ImageDataServiceUseCase{imageStoragePort: imageStoragePort}
}

func (i ImageDataServiceUseCase) SaveImageData(imageData *model.ImageData) (imageUrl string, err error) {
	imageUrl, err = i.imageStoragePort.SaveImageInStorage(imageData)

	if err != nil {
		return "", err
	}

	return imageUrl, nil
}

package usecase

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/api"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"time"
)

// UploadUseCase implements IUploadServicePort
type UploadUseCase struct {
	imageServicePort     api.IImageServicePort
	imageDataServicePort api.IImageDataServicePort
}

func NewUploadUseCase(imageServicePort api.IImageServicePort, imageDataServicePort api.IImageDataServicePort) *UploadUseCase {
	return &UploadUseCase{imageServicePort: imageServicePort, imageDataServicePort: imageDataServicePort}
}

func (u UploadUseCase) UploadFile(imageData *model.ImageData, contentType string) (imageUrl string, err error) {
	imageUrl, savedImageData, err := u.imageDataServicePort.SaveImageData(imageData)

	if err != nil {
		return "", err
	}

	image := model.NewImage(
		savedImageData.ID,
		savedImageData.FileName,
		imageUrl,
		int64(len(savedImageData.Data)),
		contentType,
		false,
		time.Now(),
		time.Now(),
	)

	err = u.imageServicePort.SaveImage(image)

	if err != nil {
		return "", err
	}

	return imageUrl, nil
}

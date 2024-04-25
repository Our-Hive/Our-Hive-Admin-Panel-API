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

func (u UploadUseCase) UploadFile(imageData *model.ImageData) (imageUrl string, err error) {
	imageUrl, savedImageData, err := u.imageDataServicePort.SaveImageData(imageData)

	if err != nil {
		return "", err
	}

	image := model.NewImage(
		savedImageData.ID,
		savedImageData.FileName,
		imageUrl,
		int64(len(savedImageData.Data)),
		"",
		time.Now(),
		time.Now(),
	)

	err = u.imageServicePort.SaveImage(image)

	if err != nil {
		return "", err
	}

	return imageUrl, nil
}

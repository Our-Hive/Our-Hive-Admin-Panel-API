package storageadapter

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/fbstorage/bucket"
)

// ImageStorageAdapter implements IImageStoragePort
type ImageStorageAdapter struct {
	imageBucket *bucket.ImageBucket
}

func NewImageStorageAdapter(imageBucket *bucket.ImageBucket) *ImageStorageAdapter {
	return &ImageStorageAdapter{imageBucket: imageBucket}
}

func (i ImageStorageAdapter) SaveImageInStorage(image *model.ImageData) (fileUrl string, err error) {
	url, err := i.imageBucket.SaveImageToFBStorage(image)

	if err != nil {
		return "", err
	}

	return url, nil
}

package storageadapter

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/fbstorage/bucket"
)

// ImageStorageAdapter implements IIImageStoragePort
type ImageStorageAdapter struct {
	imageBucket bucket.ImageBucket
}

func (i ImageStorageAdapter) SaveImageInStorage(image *model.ImageData) (fileUrl string, err error) {
	url, err := i.imageBucket.SaveImageToFBStorage(image)

	if err != nil {
		return "", err
	}

	return url, nil
}

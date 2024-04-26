package bucket

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/fbstorage/storageconstant"
	"github.com/google/uuid"
	"log"
	"strings"
)

type ImageBucket struct {
	bucket *storage.BucketHandle
	ctx    context.Context
}

func NewImageBucket(bucket *storage.BucketHandle, ctx context.Context) *ImageBucket {
	return &ImageBucket{bucket: bucket, ctx: ctx}
}

func (ib *ImageBucket) SaveImageToFBStorage(imageData *model.ImageData) (url string, err error) {
	objectHandle := ib.bucket.Object(imageData.FileName)

	writer := objectHandle.NewWriter(ib.ctx)
	id := uuid.New()

	imageData.ID = id.String()

	writer.ObjectAttrs.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": id.String(),
	}
	defer func(writer *storage.Writer) {
		err = writer.Close()
		if err != nil {
			log.Println("Error closing writer: ", err)
		}
	}(writer)

	if _, err = writer.Write(imageData.Data); err != nil {
		return "", err
	}

	return ib.GetImageUrl(imageData.FileName, id.String()), nil
}

func (ib *ImageBucket) GetImageUrl(filename string, id string) string {
	// if filename has a space, replace it with %20
	filename = strings.ReplaceAll(filename, " ", "%20")
	return fmt.Sprintf(storageconstant.BucketUrl, filename, id)
}

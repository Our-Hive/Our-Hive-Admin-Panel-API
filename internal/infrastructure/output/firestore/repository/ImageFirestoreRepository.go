package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/firestore/firestoreconstant"
)

type ImageFirestoreRepository struct {
	client     *firestore.Client
	ctx        context.Context
	collection string
}

func NewImageFireStoreRepository(client *firestore.Client, ctx context.Context) *ImageFirestoreRepository {
	return &ImageFirestoreRepository{client: client, ctx: ctx, collection: firestoreconstant.ImageDocumentName}
}

func (i ImageFirestoreRepository) SaveImageInCollection(image *model.Image) error {
	_, _, err := i.client.Collection(i.collection).Add(i.ctx, image)

	if err != nil {
		return err
	}

	return nil
}

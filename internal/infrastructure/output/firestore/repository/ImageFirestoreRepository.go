package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/firestore/firestoreconstant"
	"google.golang.org/api/iterator"
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
	_, err := i.client.Collection(i.collection).Doc(image.ID).Set(i.ctx, image)

	if err != nil {
		return err
	}

	return nil
}

func (i ImageFirestoreRepository) GetImageFromCollectionByName(fileName string) (*model.Image, error) {
	iter := i.client.Collection(i.collection).Where("Name", "==", fileName).Documents(i.ctx)

	for {
		doc, err := iter.Next()

		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, err
		}

		image := &model.Image{}
		err = doc.DataTo(image)

		if err != nil {
			return nil, err
		}

		return image, nil
	}

	return nil, nil
}

func (i ImageFirestoreRepository) GetAllImagesFromCollection(pageSize int, startAfter string) ([]*model.Image, error) {
	var images []*model.Image
	var query firestore.Query

	if startAfter != "" {
		doc, err := i.client.Collection(i.collection).Doc(startAfter).Get(i.ctx)

		if err != nil {
			return nil, err
		}

		query = i.client.Collection(i.collection).OrderBy("UpdatedTime", firestore.Desc).StartAfter(doc.Data()["Name"]).Limit(pageSize)
	} else {
		query = i.client.Collection(i.collection).OrderBy("UpdatedTime", firestore.Desc).Limit(pageSize)
	}

	iter := query.Documents(i.ctx)

	for {
		doc, err := iter.Next()

		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, err
		}

		image := &model.Image{}
		err = doc.DataTo(image)

		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}

func (i ImageFirestoreRepository) GetImageFromCollectionById(id string) (*model.Image, error) {
	doc, err := i.client.Collection(i.collection).Doc(id).Get(i.ctx)

	if err != nil {
		return nil, err
	}

	image := &model.Image{}
	err = doc.DataTo(image)

	if err != nil {
		return nil, err
	}

	return image, nil
}

func (i ImageFirestoreRepository) GetImagesByApprovedStatus(approved bool) ([]*model.Image, error) {
	var images []*model.Image
	iter := i.client.Collection(i.collection).Where("IsApproved", "==", approved).Documents(i.ctx)

	for {
		doc, err := iter.Next()

		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, err
		}

		image := &model.Image{}
		err = doc.DataTo(image)

		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}

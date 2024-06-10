package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
)

type ContactLineRepository struct {
	client     *firestore.Client
	ctx        context.Context
	collection string
}

func NewContactLineRepository(client *firestore.Client, ctx context.Context, collection string) *ContactLineRepository {
	return &ContactLineRepository{client: client, ctx: ctx, collection: collection}
}

func (c ContactLineRepository) SaveContactLineInCollection(contactLine *model.ContactLine) error {
	_, err := c.client.Collection(c.collection).NewDoc().Set(c.ctx, contactLine)

	if err != nil {
		return err
	}

	return nil
}

func (c ContactLineRepository) GetContactLineFromCollectionByName(name string) (*model.ContactLine, error) {
	iter := c.client.Collection(c.collection).Where("Name", "==", name).Documents(c.ctx)

	doc, err := iter.Next()

	if err != nil {
		return nil, err
	}

	var line model.ContactLine

	err = doc.DataTo(&line)

	if err != nil {
		return nil, err
	}

	return &line, nil
}

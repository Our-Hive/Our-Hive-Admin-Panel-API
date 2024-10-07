package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/firestore/firestoreconstant"
	"google.golang.org/api/iterator"
)

type ContactLineRepository struct {
	client     *firestore.Client
	ctx        context.Context
	collection string
}

func NewContactLineRepository(client *firestore.Client, ctx context.Context) *ContactLineRepository {
	return &ContactLineRepository{client: client, ctx: ctx, collection: firestoreconstant.ContactLineDocumentName}
}

func (c ContactLineRepository) SaveContactLineInCollection(contactLine *model.ContactLine) error {
	docRef := c.client.Collection(c.collection).NewDoc()
	contactLine.ID = docRef.ID
	_, err := docRef.Set(c.ctx, contactLine)

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

func (c ContactLineRepository) GetAllContactLinesFromCollection(pageSize int, startAfter string) ([]*model.ContactLine, error) {
	var contactLines []*model.ContactLine
	var query firestore.Query

	if startAfter != "" {
		doc, err := c.client.Collection(c.collection).Doc(startAfter).Get(c.ctx)

		if err != nil {
			return nil, err
		}

		query = c.client.Collection(c.collection).OrderBy("Name", firestore.Desc).StartAfter(doc.Data()).Limit(pageSize)
	} else {
		query = c.client.Collection(c.collection).OrderBy("Name", firestore.Desc).Limit(pageSize)
	}

	iter := query.Documents(c.ctx)

	for {
		doc, err := iter.Next()

		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, err
		}

		var contactLine model.ContactLine

		err = doc.DataTo(&contactLine)

		if err != nil {
			return nil, err
		}

		contactLines = append(contactLines, &contactLine)
	}

	return contactLines, nil
}

func (c ContactLineRepository) GetContactLineFromCollectionByID(id string) (*model.ContactLine, error) {
	doc, err := c.client.Collection(c.collection).Doc(id).Get(c.ctx)

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

func (c ContactLineRepository) UpdateContactLineInCollection(contactLine *model.ContactLine) error {
	_, err := c.client.Collection(c.collection).Doc(contactLine.ID).Set(c.ctx, contactLine)

	if err != nil {
		return err
	}

	return nil
}

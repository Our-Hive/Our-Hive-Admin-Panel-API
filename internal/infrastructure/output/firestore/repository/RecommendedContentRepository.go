package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/firestore/firestoreconstant"
	"google.golang.org/api/iterator"
)

type RecommendedContentRepository struct {
	client     *firestore.Client
	ctx        context.Context
	collection string
}

func NewRecommendedContentRepository(client *firestore.Client, ctx context.Context) *RecommendedContentRepository {
	return &RecommendedContentRepository{client: client, ctx: ctx, collection: firestoreconstant.RecommendedContentDocumentName}
}

func (r RecommendedContentRepository) SaveRecommendedContentInCollection(recommendedContent *model.DigitalContent) error {
	docRef := r.client.Collection(r.collection).NewDoc()
	recommendedContent.ID = docRef.ID
	_, err := docRef.Set(r.ctx, recommendedContent)

	if err != nil {
		return err
	}

	return nil
}

func (r RecommendedContentRepository) GetRecommendedContentFromCollectionByTitle(title string) (*model.DigitalContent, error) {
	iter := r.client.Collection(r.collection).Where("Title", "==", title).Documents(r.ctx)

	doc, err := iter.Next()

	if err != nil {
		return nil, err
	}

	var content model.DigitalContent

	err = doc.DataTo(&content)

	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (r RecommendedContentRepository) GetAllRecommendedContentFromCollection(pageSize int, startAfter string) ([]*model.DigitalContent, error) {
	var recommendedContents []*model.DigitalContent
	var query firestore.Query

	if startAfter != "" {
		doc, err := r.client.Collection(r.collection).Doc(startAfter).Get(r.ctx)

		if err != nil {
			return nil, err
		}

		query = r.client.Collection(r.collection).OrderBy("Title", firestore.Desc).StartAfter(doc.Data()).Limit(pageSize)
	} else {
		query = r.client.Collection(r.collection).OrderBy("Title", firestore.Desc).Limit(pageSize)
	}

	iter := query.Limit(pageSize).Documents(r.ctx)

	for {
		doc, err := iter.Next()

		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, err
		}

		var recommendedContent model.DigitalContent

		err = doc.DataTo(&recommendedContent)

		if err != nil {
			return nil, err
		}

		recommendedContents = append(recommendedContents, &recommendedContent)
	}

	return recommendedContents, nil
}

package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/firestore/firestoreconstant"
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

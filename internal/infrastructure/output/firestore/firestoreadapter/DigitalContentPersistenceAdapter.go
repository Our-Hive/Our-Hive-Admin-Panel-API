package firestoreadapter

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/firestore/repository"
)

// DigitalContentPersistenceAdapter implements IDigitalContentPersistencePort
type DigitalContentPersistenceAdapter struct {
	recommendedContentRepository *repository.RecommendedContentRepository
}

func NewDigitalContentPersistenceAdapter(recommendedContentRepository *repository.RecommendedContentRepository) *DigitalContentPersistenceAdapter {
	return &DigitalContentPersistenceAdapter{recommendedContentRepository: recommendedContentRepository}
}

func (d DigitalContentPersistenceAdapter) SaveDigitalContentInDatabase(content *model.DigitalContent) (err error) {
	err = d.recommendedContentRepository.SaveRecommendedContentInCollection(content)

	if err != nil {
		return err
	}

	return nil
}

func (d DigitalContentPersistenceAdapter) GetDigitalContentFromDatabaseByTitle(title string) (*model.DigitalContent, error) {
	content, err := d.recommendedContentRepository.GetRecommendedContentFromCollectionByTitle(title)

	if err != nil {
		return nil, err
	}

	return content, nil
}

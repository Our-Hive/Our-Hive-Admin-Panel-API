package firestoreadapter

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/firestore/repository"
)

type ImagePersistenceAdapter struct {
	imageRepository repository.ImageFirestoreRepository
}

func (i ImagePersistenceAdapter) SaveImageInDatabase(image *model.Image) error {
	err := i.imageRepository.SaveImageInCollection(image)

	if err != nil {
		return err
	}

	return nil
}

package firestoreadapter

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/firestore/repository"
)

// ImagePersistenceAdapter implements IImagePersistencePort
type ImagePersistenceAdapter struct {
	imageRepository *repository.ImageFirestoreRepository
}

func NewImagePersistenceAdapter(imageRepository *repository.ImageFirestoreRepository) *ImagePersistenceAdapter {
	return &ImagePersistenceAdapter{imageRepository: imageRepository}
}

func (i ImagePersistenceAdapter) SaveImageInDatabase(image *model.Image) error {
	err := i.imageRepository.SaveImageInCollection(image)

	if err != nil {
		return err
	}

	return nil
}

func (i ImagePersistenceAdapter) GetImageFromDatabaseByName(fileName string) (*model.Image, error) {
	image, err := i.imageRepository.GetImageFromCollectionByName(fileName)

	if err != nil {
		return nil, err
	}

	return image, nil
}

func (i ImagePersistenceAdapter) GetAllImagesFromDatabase(pageSize int, startAfter string) ([]*model.Image, error) {
	images, err := i.imageRepository.GetAllImagesFromCollection(pageSize, startAfter)

	if err != nil {
		return nil, err
	}

	return images, nil
}

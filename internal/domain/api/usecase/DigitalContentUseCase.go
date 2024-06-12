package usecase

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/domainerror"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/spi"
)

// DigitalContentUseCase implements IDigitalContentServicePort
type DigitalContentUseCase struct {
	digitalContentPersistencePort spi.IDigitalContentPersistencePort
}

func NewDigitalContentUseCase(digitalContentPersistencePort spi.IDigitalContentPersistencePort) *DigitalContentUseCase {
	return &DigitalContentUseCase{digitalContentPersistencePort: digitalContentPersistencePort}
}

func (d DigitalContentUseCase) CreateDigitalContent(content *model.DigitalContent) (err error) {
	foundContent, _ := d.digitalContentPersistencePort.GetDigitalContentFromDatabaseByTitle(content.Title)

	if foundContent != nil {
		return &domainerror.DigitalContentAlreadyExistsError{Title: content.Title}
	}

	err = d.digitalContentPersistencePort.SaveDigitalContentInDatabase(content)

	if err != nil {
		return err
	}

	return nil
}

func (d DigitalContentUseCase) GetAllDigitalContent(pageSize int, startAfter string) (contents []*model.DigitalContent, err error) {
	contents, err = d.digitalContentPersistencePort.GetAllDigitalContentFromDatabase(pageSize, startAfter)

	if err != nil {
		return nil, err
	}

	return contents, nil
}

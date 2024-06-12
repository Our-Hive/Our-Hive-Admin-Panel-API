package spi

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IDigitalContentPersistencePort interface {
	SaveDigitalContentInDatabase(content *model.DigitalContent) (err error)
	GetDigitalContentFromDatabaseByTitle(title string) (*model.DigitalContent, error)
}

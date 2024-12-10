package spi

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IContactLinePersistencePort interface {
	SaveContactLineInDatabase(line *model.ContactLine) (err error)
	GetContactLineFromDatabaseByName(name string) (*model.ContactLine, error)
	GetAllContactLinesFromDatabase(pageSize int, startAfter string) ([]*model.ContactLine, error)
	GetContactLineFromDatabaseByID(id string) (*model.ContactLine, error)
	UpdateContactLineInDatabase(line *model.ContactLine) (err error)
}

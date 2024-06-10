package spi

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IContactLinePersistencePort interface {
	SaveContactLineInDatabase(line *model.ContactLine) (err error)
}
package api

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IContactLineServicePort interface {
	CreateContactLine(line *model.ContactLine) (err error)
}

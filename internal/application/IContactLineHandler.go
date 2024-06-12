package application

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
)

type IContactLineHandler interface {
	Create(request *request.CreateContactLine) (httpStatus int, err error)
	GetAll(pageSize int, startAfter string) (contactLines []*model.ContactLine, httpStatus int, err error)
}

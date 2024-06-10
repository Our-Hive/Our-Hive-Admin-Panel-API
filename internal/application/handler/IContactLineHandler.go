package handler

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
)

type IContactLineHandler interface {
	Create(request *request.CreateContactLine) (httpStatus int, err error)
}

package application

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/response"
)

type IImageHandler interface {
	GetAll(pageSize int, startAfter string) (images []*response.Image, httpStatus int, err error)
	Approve(id string) (httpStatus int, err error)
	GetByApprovedStatus(approved bool, pageSize int, startAfter string) (images []*response.Image, httpStatus int, err error)
}

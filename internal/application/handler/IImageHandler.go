package handler

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IImageHandler interface {
	GetAll(pageSize int, startAfter string) (images []*model.Image, httpStatus int, err error)
}

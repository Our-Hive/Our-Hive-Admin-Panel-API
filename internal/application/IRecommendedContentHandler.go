package application

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
)

type IRecommendedContentHandler interface {
	Create(request *request.CreateDigitalContent) (httpStatus int, err error)
	GetAll() (content []*model.DigitalContent, httpStatus int, err error)
}

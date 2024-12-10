package application

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/response"
)

type IRecommendedContentHandler interface {
	Create(request *request.CreateDigitalContent) (httpStatus int, err error)
	GetAll() (content []*response.RecommendedContent, httpStatus int, err error)
}

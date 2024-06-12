package application

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"

type IRecommendedContentHandler interface {
	Create(request *request.CreateDigitalContent) (httpStatus int, err error)
}

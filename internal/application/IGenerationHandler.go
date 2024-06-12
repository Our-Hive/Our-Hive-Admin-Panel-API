package application

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/response"
)

type IGenerationHandler interface {
	GenerateIAImage(request *request.GenerateIAImage) (response *response.UploadImage, httpStatus int, err error)
}

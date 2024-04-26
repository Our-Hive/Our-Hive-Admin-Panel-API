package handler

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/response"

type IUploadHandler interface {
	Upload(file []byte, fileName string, contentType string) (responseBody *response.UploadImage, httpStatus int, err error)
}

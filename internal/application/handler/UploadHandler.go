package handler

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/response"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/api"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"net/http"
	"path/filepath"
)

// UploadHandler implements IUploadHandler
type UploadHandler struct {
	uploadServicePort api.IUploadServicePort
}

func NewUploadHandler(uploadServicePort api.IUploadServicePort) *UploadHandler {
	return &UploadHandler{uploadServicePort: uploadServicePort}
}

func (u UploadHandler) Upload(file []byte, fileName string, contentType string) (responseBody *response.UploadImage, httpStatus int, err error) {
	formattedFileName := removeExtension(fileName)

	imageData := &model.ImageData{
		Data:     file,
		FileName: formattedFileName,
	}

	imageUrl, err := u.uploadServicePort.UploadFile(imageData, contentType)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	responseBody = &response.UploadImage{
		Url: imageUrl,
	}

	return responseBody, 200, nil
}

func removeExtension(filename string) string {
	extension := filepath.Ext(filename)
	return filename[0 : len(filename)-len(extension)]
}

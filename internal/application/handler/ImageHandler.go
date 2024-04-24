package handler

import (
	"errors"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/api"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/domainerror"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"net/http"
)

// ImageHandler implements IImageHandler
type ImageHandler struct {
	imageServicePort api.IImageServicePort
}

func NewImageHandler(imageServicePort api.IImageServicePort) *ImageHandler {
	return &ImageHandler{imageServicePort: imageServicePort}
}

func (i ImageHandler) GetAll(pageSize int, startAfter string) (images []*model.Image, httpStatus int, err error) {
	images, err = i.imageServicePort.GetAllImages(pageSize, startAfter)

	var noDataFoundError *domainerror.NoDataFound
	if errors.As(err, &noDataFoundError) {
		return nil, http.StatusNotFound, noDataFoundError
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return images, http.StatusOK, nil
}

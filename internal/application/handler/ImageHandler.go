package handler

import (
	"errors"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/response"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/mapper"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/api"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/domainerror"
	"net/http"
)

// ImageHandler implements IImageHandler
type ImageHandler struct {
	imageServicePort api.IImageServicePort
}

func NewImageHandler(imageServicePort api.IImageServicePort) *ImageHandler {
	return &ImageHandler{imageServicePort: imageServicePort}
}

func (i ImageHandler) GetAll(pageSize int, startAfter string) (images []*response.Image, httpStatus int, err error) {
	retrievedImages, err := i.imageServicePort.GetAllImages(pageSize, startAfter)

	var noDataFoundError *domainerror.NoDataFound
	if errors.As(err, &noDataFoundError) {
		return nil, http.StatusNotFound, noDataFoundError
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	images = mapper.FromImagesToImageResponses(retrievedImages)

	return images, http.StatusOK, nil
}

func (i ImageHandler) Approve(id string) (httpStatus int, err error) {
	err = i.imageServicePort.ApproveImage(id)

	var noDataFoundError *domainerror.NoDataFound
	if errors.As(err, &noDataFoundError) {
		return http.StatusNotFound, noDataFoundError
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}

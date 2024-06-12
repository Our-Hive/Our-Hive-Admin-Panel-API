package handler

import (
	"errors"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/api"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// RecommendedContentHandler implements IRecommendedContentHandler
type RecommendedContentHandler struct {
	digitalContentServicePort api.IDigitalContentServicePort
}

func NewRecommendedContentHandler(digitalContentServicePort api.IDigitalContentServicePort) *RecommendedContentHandler {
	return &RecommendedContentHandler{digitalContentServicePort: digitalContentServicePort}
}

func (r RecommendedContentHandler) Create(request *request.CreateDigitalContent) (httpStatus int, err error) {
	validate := validator.New()

	err = validate.Struct(request)

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		return http.StatusBadRequest, validationErrors
	}

	digitalContent := model.NewDigitalContent(request.Title, request.Description, request.URL)

	err = r.digitalContentServicePort.CreateDigitalContent(digitalContent)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (r RecommendedContentHandler) GetAll(pageSize int, startAfter string) (contents []*model.DigitalContent, httpStatus int, err error) {
	contents, err = r.digitalContentServicePort.GetAllDigitalContent(pageSize, startAfter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return contents, http.StatusOK, nil
}

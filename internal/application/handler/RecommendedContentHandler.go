package handler

import (
	"errors"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/response"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/mapper"
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

func (r RecommendedContentHandler) GetAll() (content []*response.RecommendedContent, httpStatus int, err error) {
	digitalContent, err := r.digitalContentServicePort.GetAllDigitalContent()

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return mapper.FromRecommendedContentModelsToResponse(digitalContent), http.StatusOK, nil
}

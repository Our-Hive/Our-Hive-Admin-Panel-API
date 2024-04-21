package handler

import (
	"errors"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/response"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/api"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/domainerror"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// GenerationHandler implements IGenerationHandler
type GenerationHandler struct {
	generationUseCase api.IGenerationServicePort
}

func NewGenerationHandler(generationUseCase api.IGenerationServicePort) *GenerationHandler {
	return &GenerationHandler{generationUseCase: generationUseCase}
}

func (g GenerationHandler) GenerateIAImage(request *request.GenerateIAImage) (resp *response.GenerateIAImage, httpStatus int, err error) {
	validate := validator.New()

	err = validate.Struct(request)

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		return nil, http.StatusBadRequest, validationErrors
	}

	url, err := g.generationUseCase.GenerateImage(request.Prompt, request.FileName)

	var imageAlreadyExistsError *domainerror.ImageAlreadyExistsError
	if errors.As(err, &imageAlreadyExistsError) {
		return nil, http.StatusConflict, imageAlreadyExistsError
	}

	var notEthicalPromptError *domainerror.IsNotEthicalPromptError
	if errors.As(err, &notEthicalPromptError) {
		return nil, http.StatusUnprocessableEntity, notEthicalPromptError
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	responseBody := &response.GenerateIAImage{Url: url}

	return responseBody, http.StatusCreated, nil
}

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

func (g GenerationHandler) GenerateIAImage(request *request.GenerateIAImage) (response *response.GenerateIAImage, httpStatus int, err error) {
	validate := validator.New()

	err = validate.Struct(request)

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		return nil, http.StatusBadRequest, validationErrors
	}

	url, err := g.generationUseCase.GenerateImage(request.Prompt, request.FileName)

	if errors.Is(err, &domainerror.IsNotEthicalPromptError{}) {
		return nil, http.StatusUnprocessableEntity, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	response.Url = url

	return response, http.StatusOK, nil
}

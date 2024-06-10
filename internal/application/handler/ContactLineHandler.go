package handler

import (
	"errors"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/api"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/domainerror"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ContactLineHandler implements IContactLineHandler
type ContactLineHandler struct {
	contactLineServicePort api.IContactLineServicePort
}

func NewContactLineHandler(contactLineServicePort api.IContactLineServicePort) *ContactLineHandler {
	return &ContactLineHandler{contactLineServicePort: contactLineServicePort}
}

func (c ContactLineHandler) Create(request *request.CreateContactLine) (httpStatus int, err error) {
	validate := validator.New()

	err = validate.Struct(request)

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		return http.StatusBadRequest, validationErrors
	}

	contactLine := &model.ContactLine{
		Name:        request.Name,
		Description: request.Description,
	}
	err = c.contactLineServicePort.CreateContactLine(contactLine)

	var contactLineAlreadyExistsError *domainerror.ContactLineAlreadyExistsError
	if errors.As(err, &contactLineAlreadyExistsError) {
		return http.StatusConflict, contactLineAlreadyExistsError
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (c ContactLineHandler) GetAll(pageSize int, startAfter string) (contactLines []*model.ContactLine, httpStatus int, err error) {
	contactLines, err = c.contactLineServicePort.GetAllContactLines(pageSize, startAfter)

	var noDataFoundError *domainerror.NoDataFound
	if errors.As(err, &noDataFoundError) {
		return nil, http.StatusNotFound, noDataFoundError
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return contactLines, http.StatusOK, nil
}

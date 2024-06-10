package usecase

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/domainerror"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/spi"
)

// ContactLineUseCase implements IContactLineServicePort
type ContactLineUseCase struct {
	contactLinePersistencePort spi.IContactLinePersistencePort
}

func NewContactLineUseCase(contactLinePersistencePort spi.IContactLinePersistencePort) *ContactLineUseCase {
	return &ContactLineUseCase{contactLinePersistencePort: contactLinePersistencePort}
}

func (c ContactLineUseCase) CreateContactLine(line *model.ContactLine) (err error) {
	foundLine, _ := c.GetContactLineByName(line.Name)

	if foundLine != nil {
		return &domainerror.ContactLineAlreadyExistsError{Name: line.Name}
	}

	err = c.contactLinePersistencePort.SaveContactLineInDatabase(line)

	if err != nil {
		return err
	}

	return nil
}

func (c ContactLineUseCase) GetContactLineByName(name string) (*model.ContactLine, error) {
	line, err := c.contactLinePersistencePort.GetContactLineFromDatabaseByName(name)

	if err != nil {
		return nil, err
	}

	if line == nil {
		return nil, &domainerror.ContactLineNotFoundError{Name: name}
	}

	return line, nil
}

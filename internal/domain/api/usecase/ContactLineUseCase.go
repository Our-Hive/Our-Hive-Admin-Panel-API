package usecase

import (
	"fmt"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/domainconstant"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/domainerror"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/spi"
	"strings"
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

func (c ContactLineUseCase) GetAllContactLines(pageSize int, startAfter string) ([]*model.ContactLine, error) {
	lines, err := c.contactLinePersistencePort.GetAllContactLinesFromDatabase(pageSize, startAfter)

	if err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return nil, &domainerror.NoDataFound{Message: fmt.Sprintf(domainconstant.NoDataFoundErrorMessage, "contact lines")}
	}

	return lines, nil
}

func (c ContactLineUseCase) UpdateContactLine(line *model.ContactLine) (err error) {
	foundLine, _ := c.contactLinePersistencePort.GetContactLineFromDatabaseByID(line.ID)

	if foundLine == nil {
		return &domainerror.ContactLineNotFoundError{
			Name: line.ID,
		}
	}

	updateContactLineFields(foundLine, line)

	err = c.contactLinePersistencePort.UpdateContactLineInDatabase(foundLine)

	if err != nil {
		return err
	}

	return nil
}

func updateContactLineFields(foundLine, line *model.ContactLine) {
	if strings.TrimSpace(line.Name) != "" {
		foundLine.Name = line.Name
	}

	if strings.TrimSpace(line.Description) != "" {
		foundLine.Description = line.Description
	}
}

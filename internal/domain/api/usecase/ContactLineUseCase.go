package usecase

import (
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
	err = c.contactLinePersistencePort.SaveContactLineInDatabase(line)

	if err != nil {
		return err
	}

	return nil
}

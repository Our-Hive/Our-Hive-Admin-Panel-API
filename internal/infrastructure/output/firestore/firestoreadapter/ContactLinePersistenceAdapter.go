package firestoreadapter

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/firestore/repository"
)

// ContactLinePersistenceAdapter implements IContactLinePersistencePort
type ContactLinePersistenceAdapter struct {
	contactLineRepository *repository.ContactLineRepository
}

func NewContactLinePersistenceAdapter(contactLineRepository *repository.ContactLineRepository) *ContactLinePersistenceAdapter {
	return &ContactLinePersistenceAdapter{contactLineRepository: contactLineRepository}
}

func (c ContactLinePersistenceAdapter) SaveContactLineInDatabase(line *model.ContactLine) error {
	err := c.contactLineRepository.SaveContactLineInCollection(line)

	if err != nil {
		return err
	}

	return nil
}

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

func (c ContactLinePersistenceAdapter) GetContactLineFromDatabaseByName(name string) (*model.ContactLine, error) {
	line, err := c.contactLineRepository.GetContactLineFromCollectionByName(name)

	if err != nil {
		return nil, err
	}

	return line, nil
}

func (c ContactLinePersistenceAdapter) GetAllContactLinesFromDatabase(pageSize int, startAfter string) ([]*model.ContactLine, error) {
	lines, err := c.contactLineRepository.GetAllContactLinesFromCollection(pageSize, startAfter)

	if err != nil {
		return nil, err
	}

	return lines, nil
}

package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/entities"
	mockdatabase "github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/pkg/mock_database"
)

type InventryService struct {
	inventry *mockdatabase.Inventry
}

func NewInventryService(inventry *mockdatabase.Inventry) *InventryService {
	return &InventryService{
		inventry: inventry,
	}
}
func (i InventryService) Stocks() ([]entities.Item, error) {
	items := make([]entities.Item, 0)
	item := entities.Item{
		ID:    i.inventry.Item1.ID,
		Name:  i.inventry.Item1.Name,
		Count: i.inventry.Item1.Count,
	}

	items = append(items, item)
	item = entities.Item{
		ID:    i.inventry.Item2.ID,
		Name:  i.inventry.Item2.Name,
		Count: i.inventry.Item2.Count,
	}
	items = append(items, item)

	return items, nil
}

func (i InventryService) GetByID(id uuid.UUID) (*entities.Item, error) {
	if i.inventry.Item1.ID == id {
		return &entities.Item{
			ID:    i.inventry.Item1.ID,
			Name:  i.inventry.Item1.Name,
			Count: i.inventry.Item1.Count,
		}, nil
	}

	if i.inventry.Item2.ID == id {
		return &entities.Item{
			ID:    i.inventry.Item2.ID,
			Name:  i.inventry.Item2.Name,
			Count: i.inventry.Item2.Count,
		}, nil
	}

	return nil, errors.New("not found")
}

func (i InventryService) BuyByID(id uuid.UUID, count int64) error {
	if i.inventry.Item1.ID == id && i.inventry.Item1.Count > count {
		i.inventry.Item1.Count -= count
	} else if i.inventry.Item2.ID == id && i.inventry.Item2.Count > count {
		i.inventry.Item2.Count -= count
	}

	return errors.New("not found")
}

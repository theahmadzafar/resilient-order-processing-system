package mockdatabase

import "github.com/google/uuid"

func NewMockConnection() (*Inventry, error) {
	inv := Inventry{
		item1: InventryItem{ID: uuid.New(), Name: "item1", Count: 10},
		item2: InventryItem{ID: uuid.New(), Name: "item2", Count: 10},
	}

	return &inv, nil
}

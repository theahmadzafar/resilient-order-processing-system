package mockdatabase

import "github.com/google/uuid"

func NewMockInventry() (*Inventry, error) {
	inv := Inventry{
		Item1: InventryItem{ID: uuid.New(), Name: "item1", Count: 10},
		Item2: InventryItem{ID: uuid.New(), Name: "item2", Count: 10},
	}

	return &inv, nil
}

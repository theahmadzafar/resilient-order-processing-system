package mockdatabase

import "github.com/google/uuid"

type InventryItem struct {
	// item id
	ID uuid.UUID
	// name from item
	Name string
	// available count
	Count int64
}

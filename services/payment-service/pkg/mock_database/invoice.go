package mockdatabase

import "github.com/google/uuid"

type Invoice struct {
	// item id
	ID uuid.UUID
	// item id
	OrderID uuid.UUID
}

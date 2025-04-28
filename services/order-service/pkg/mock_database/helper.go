package mockdatabase

import "github.com/google/uuid"

type Item struct {
	// item id
	ID uuid.UUID
	// name from item
	Name string
	// available count
	Count int64
}

type Order struct {
	// item id
	ID uuid.UUID
	// name from item
	Status string
	// available count
	Items []Item
	// paymentID
	PaymentID string
}

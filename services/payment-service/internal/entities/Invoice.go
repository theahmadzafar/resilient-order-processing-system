package entities

import "github.com/google/uuid"

type PaymentRequest struct {
	// item id
	ID uuid.UUID
	// item id
	OrderID uuid.UUID
}

package entities

import "github.com/google/uuid"

type ItemOrderRequest struct {
	ID    uuid.UUID `json:"id" required:"true"`
	Count int64     `json:"count" required:"true"`
}

type OrderRequest struct {
	List []ItemOrderRequest `json:"list" required:"true"`
}

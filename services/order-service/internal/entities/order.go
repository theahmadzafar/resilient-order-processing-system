package entities

import "github.com/google/uuid"

type ItemOrderRequest struct {
	ID    uuid.UUID `json:"id" required:"true"`
	Count int64     `json:"count" required:"true"`
}

type OrderRequest struct {
	List []ItemOrderRequest `json:"list" required:"true"`
}

type GetOrderRequest struct {
	OrderID uuid.UUID `json:"order_id"  form:"order_id" required:"true"`
}

type GetOrderResponse struct {
	// item id
	ID uuid.UUID `json:"id"`
	// name from item
	Status string `json:"status"`
	// available count
	Items []Item `json:"items"`
	// paymentID
	PaymentID uuid.UUID `json:"payment_id"`
}

type Item struct {
	// item id
	ID uuid.UUID `json:"id"`
	// name from item
	Name string `json:"name"`
	// available count
	Count int64 `json:"count"`
}

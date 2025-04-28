package entities

import "github.com/google/uuid"

type Item struct {
	// item id
	ID uuid.UUID `json:"Id"`
	// name from item
	Name string `json:"name"`
	// available count
	Count int64 `json:"count"`
}

package entities

import "github.com/google/uuid"

type Item struct {
	// item id
	ID uuid.UUID
	// name from item
	Name string
}

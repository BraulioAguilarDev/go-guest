package entity

import "github.com/google/uuid"

// ID entity
type ID = uuid.UUID

// NewID func
func NewID() ID {
	return ID(uuid.New())
}

// StringToID func
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

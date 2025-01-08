package store

import "github.com/google/uuid"

// A Category is a logical grouping of asset types
type Category struct {
	Id          uuid.UUID
	Title       string
	Description *string
}

package store

import "github.com/google/uuid"

// An Asset is representative of an entity that requires maintenance

// TODO: add references to categories and groups
type Asset struct {
	Id           uuid.UUID
	Title        string
	Year         *int
	Manufacturer *string
	Make         *string
	ModelNumber  *string
	SerialNumber *string
	Description  *string
}

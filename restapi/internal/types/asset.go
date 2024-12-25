package types

import "github.com/google/uuid"

// An Asset is representative of an entity that requires maintenance
type Asset struct {
	Id           uuid.UUID `json:"id" swaggerignore:"true"`
	Title        string    `json:"title" binding:"required"`
	Year         *int      `json:"year"`
	Make         *string   `json:"make"`
	ModelNumber  *string   `json:"modelNumber"`
	SerialNumber *string   `json:"serialNumber"`
	Description  *string   `json:"description"`
}

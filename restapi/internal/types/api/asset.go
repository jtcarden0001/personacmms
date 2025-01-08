package api

import "github.com/google/uuid"

// An Asset is representative of an entity that requires maintenance
type AssetRequest struct {
	Title        string  `json:"title" binding:"required"`
	Year         *int    `json:"year"`
	Manufacturer *string `json:"manufacturer"`
	Make         *string `json:"make"`
	ModelNumber  *string `json:"modelNumber"`
	SerialNumber *string `json:"serialNumber"`
	Description  *string `json:"description"`
}

type AssetResponse struct {
	Id           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Year         *int      `json:"year"`
	Manufacturer *string   `json:"manufacturer"`
	Make         *string   `json:"make"`
	ModelNumber  *string   `json:"modelNumber"`
	SerialNumber *string   `json:"serialNumber"`
	Description  *string   `json:"description"`
	// TODO:
	// array string of api references "/api/v1/categories/{id}"
	// Categories []string `json:"associatedCategories"`
	// array string of api references "/api/v1/groups/{id}"
	// Groups []string `json:"associatedGroups"`
}

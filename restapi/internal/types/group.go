package types

import "github.com/google/uuid"

// group is a logial grouping (in other words, a container) of related assets
type Group struct {
	Title string    `json:"title" binding:"required"`
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
}

package store

import "github.com/google/uuid"

// group is a logial grouping (in other words, a container) of related assets
type Group struct {
	Id    uuid.UUID
	Title string
}

package types

import uid "github.com/google/uuid"

type Category struct {
	Id    uid.UUID `json:"id"`
	Title string   `json:"title" binding:"required"`
}

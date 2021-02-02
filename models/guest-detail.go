package models

import "github.com/brauliodev29/go-guest/pkg/entity"

// GuestDetail struct
type GuestDetail struct {
	ID              entity.ID `json:"id"`
	Adult           int       `json:"adult"`
	AdultUpdated    int       `json:"adult_updated"`
	Children        int       `json:"children"`
	ChildrenUpdated int       `json:"children_updated"`
	Total           int       `json:"total"`
}

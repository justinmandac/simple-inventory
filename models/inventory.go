package models

// ItemCategory - an item category
type ItemCategory struct {
	Name     string `json:"name"`
	ID       int    `json:"id"`
	ParentID int    `json:"parentId"`
}

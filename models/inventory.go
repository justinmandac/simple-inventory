package models

import (
	"database/sql"
)

// ItemCategory - an item category
type ItemCategory struct {
	Name     string        `json:"name"`
	ID       int           `json:"id"`
	ParentID sql.NullInt64 `json:"parentId"`
}

// Item - an Inventory Item
type Item struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Categories  []ItemCategory `json:"categories"`
	Stock       int            `json:"stock"`
	Price       float32        `json:"price"` // Price per unit
}

// ItemStock - Stock information
type ItemStock struct {
	ID       int     `json:"id"`
	Unit     string  `json:"unit"`
	Quantity float32 `json:"quantity"`
}

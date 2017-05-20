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

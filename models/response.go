package models

// Response - generic JSON response
type Response struct {
	Err     int         `json:"err"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// ItemCategory - an item category
type ItemCategory struct {
	Name     string `json:"name"`
	ID       int    `json:"id"`
	ParentID int    `json:"parentId"`
}

package models

// Response - generic JSON response
type Response struct {
	Err     int         `json:"err"`
	Message []string    `json:"msg"`
	Data    interface{} `json:"data"`
}

package models

// Response - generic response
type Response struct {
	Err     int
	Message string
	Data    interface{}
}

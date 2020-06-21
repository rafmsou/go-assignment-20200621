package models

// User represents a user table schema in the database
type User struct {
	ID           int
	AgentID      int
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
}

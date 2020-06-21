package models

// Policy represents a schema in the database
type Policy struct {
	ID           int
	MobileNumber string  `json:"mobile_number"`
	Premium      float64 `json:"premium"`
	Type         string  `json:"type"`
}

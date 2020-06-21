package services

import (
	"github.com/rafmsou/agentero/database"
	"github.com/rafmsou/agentero/models"
)

// GetPoliciesByMobileNumber returns all policies by contact's mobile number
func GetPoliciesByMobileNumber(mobileNumber string) ([]*models.Policy, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	// Create read-only transaction
	txn := db.Txn(false)
	defer txn.Abort()

	// Get all policies for this user by using its mobile number
	it, err := txn.Get("policy", "mobile_number", mobileNumber)
	if err != nil {
		return nil, err
	}

	policies := []*models.Policy{}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.Policy)
		policies = append(policies, &models.Policy{
			ID:           p.ID,
			MobileNumber: p.MobileNumber,
			Premium:      p.Premium,
			Type:         p.Type,
		})
	}

	return policies, nil
}

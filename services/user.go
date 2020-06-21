package services

import (
	"errors"
	"fmt"

	"github.com/rafmsou/agentero/database"
	"github.com/rafmsou/agentero/models"
)

func getUser(field string, value interface{}) (*models.User, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	// Create read-only transaction
	txn := db.Txn(false)
	defer txn.Abort()

	// Get user from the database
	user, err := txn.First("user", field, value)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("Could not find user")
	}

	u := user.(models.User)
	ret := &models.User{
		ID:           u.ID,
		Name:         u.Name,
		MobileNumber: u.MobileNumber,
	}
	return ret, nil
}

// GetUserByID returns a user record from the database by its id
func GetUserByID(userID int64) (*models.User, error) {
	user, err := getUser("id", userID)
	if err != nil {
		return nil, fmt.Errorf("Could not find user with ID: %d", userID)
	}
	return user, nil
}

// GetUserByMobileNumber returns a user record from the database by its mobile_number
func GetUserByMobileNumber(mobileNumber string) (*models.User, error) {
	user, err := getUser("mobile_number", mobileNumber)
	if err != nil {
		return nil, fmt.Errorf("Could not find user with MobileNumber: %s", mobileNumber)
	}
	return user, nil
}

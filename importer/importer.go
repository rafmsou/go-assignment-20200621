package importer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/rafmsou/agentero/database"
	"github.com/rafmsou/agentero/models"
)

// RunAtInterval will run the AMS import process at each specified `interval`
func RunAtInterval(serverAddress string, agentID int, interval time.Duration) {
	time.Sleep(interval)
	Run(serverAddress, agentID)
	RunAtInterval(serverAddress, agentID, interval)
}

// Run will import data from remote location `serverAddress` for `agentID`
func Run(serverAddress string, agentID int) error {
	log.Printf("Running importer from %s for agentID %d", serverAddress, agentID)

	// Get users and policies from external service
	users, err := getUsers(serverAddress, agentID)
	if err != nil {
		return err
	}
	policies, err := getPolicies(serverAddress, agentID)
	if err != nil {
		return err
	}

	// Insert users and policies into the database
	db, err := database.GetInstance()
	if err != nil {
		return err
	}
	txn := db.Txn(true)
	for idx, u := range users {
		u.ID = idx + 1
		u.MobileNumber = cleanPhoneNumber(u.MobileNumber)
		u.AgentID = agentID
		if err := txn.Insert("user", u); err != nil {
			return err
		}
	}
	for idx, p := range policies {
		p.ID = idx + 1
		p.MobileNumber = cleanPhoneNumber(p.MobileNumber)
		if err := txn.Insert("policy", p); err != nil {
			return err
		}
	}
	txn.Commit()

	return nil
}

func cleanPhoneNumber(mobileNumber string) string {
	// Regex to match all characters except numbers and replace them with empty string
	var re = regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(mobileNumber, "")
}

func getUsers(serverAddress string, agentID int) ([]models.User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%d", serverAddress, agentID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	users := []models.User{}
	jsonErr := json.Unmarshal(body, &users)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return users, nil
}

func getPolicies(serverAddress string, agentID int) ([]models.Policy, error) {
	resp, err := http.Get(fmt.Sprintf("%s/policies/%d", serverAddress, agentID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	policies := []models.Policy{}
	jsonErr := json.Unmarshal(body, &policies)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return policies, nil
}

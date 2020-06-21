package importer

import (
	"testing"

	"github.com/jarcoal/httpmock"
	amsmock "github.com/rafmsou/agentero/ams_mock"
	"github.com/rafmsou/agentero/database"
	"github.com/rafmsou/agentero/models"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	amsmock.MockHTTPRequests()
	defer httpmock.DeactivateAndReset()

	err := Run(amsmock.ServerAddress, amsmock.AgentID)
	if err != nil {
		t.Error(err)
	}
	db, err := database.GetInstance()
	if err != nil {
		t.Error(err)
	}

	// Create read-only transaction
	txn := db.Txn(false)
	defer txn.Abort()

	// Try to get a single user from the database
	user, err := txn.First("user", "id", 1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(
		models.User{ID: 1, AgentID: 10, Name: "user1", MobileNumber: "1234567890"},
		user,
	)

	// Try to get a single policy from the database
	policy, err := txn.First("policy", "id", 1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(
		models.Policy{ID: 1, MobileNumber: "1234567890", Premium: 2000, Type: "homeowner"},
		policy,
	)
}

func TestRunWithServerDown(t *testing.T) {
	assert := assert.New(t)
	err := Run(amsmock.ServerAddress, amsmock.AgentID)
	assert.Error(err)
}

package services

import (
	"testing"

	"github.com/jarcoal/httpmock"
	amsmock "github.com/rafmsou/agentero/ams_mock"
	"github.com/rafmsou/agentero/importer"
	"github.com/stretchr/testify/assert"
)

func init() {
	httpmock.Activate()
	amsmock.MockHTTPRequests()
	defer httpmock.DeactivateAndReset()

	importer.Run(amsmock.ServerAddress, amsmock.AgentID)
}

func TestGetUserByID(t *testing.T) {
	assert := assert.New(t)

	user, err := GetUserByID(5)
	if err != nil {
		t.Error(err)
	}
	assert.Equal("user5", user.Name)
}

func TestGetUserByIDWithNoMatch(t *testing.T) {
	assert := assert.New(t)

	user, err := GetUserByID(6)
	assert.Nil(user)
	assert.Error(err)
}

func TestGetUserByMobileNumber(t *testing.T) {
	assert := assert.New(t)

	user, err := GetUserByMobileNumber("1234567893")
	if err != nil {
		t.Error(err)
	}
	assert.Equal("user4", user.Name)
}

func TestGetUserByMobileNumberWithNoMatch(t *testing.T) {
	assert := assert.New(t)

	user, err := GetUserByMobileNumber("0000000000")
	assert.Nil(user)
	assert.Error(err)
}

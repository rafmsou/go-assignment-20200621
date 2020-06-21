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

func TestGetPoliciesByMobileNumber(t *testing.T) {
	assert := assert.New(t)

	policies, err := GetPoliciesByMobileNumber("1234567890")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(2, len(policies))
}

func TestGetPoliciesByMobileNumberWithNoMatch(t *testing.T) {
	assert := assert.New(t)

	policies, err := GetPoliciesByMobileNumber("0000000000")
	assert.Empty(policies)
	assert.Nil(err)
}

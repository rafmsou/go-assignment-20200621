package amsmock

import (
	"fmt"

	"github.com/jarcoal/httpmock"
)

// ServerAddress represents a fake url to be used while calling mocked AMS endpoints to import data
const ServerAddress = "http://ams.test"

// AgentID is the agent id to be used while calling mocked AMS endpoints to import data
const AgentID = 10

// MockHTTPRequests set up a mock to return fake data from AMS HTTP endpoints
func MockHTTPRequests() {
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/users/%d", ServerAddress, AgentID),
		httpmock.NewStringResponder(
			200,
			`[
				{"name": "user1", "mobile_number": "1234567890"},
				{"name": "user2", "mobile_number": "123 456 7891"},
				{"name": "user3", "mobile_number": "(123) 456 7892"},
				{"name": "user4", "mobile_number": "(123) 456-7893"},
				{"name": "user5", "mobile_number": "123-456-7894"}
			]`,
		),
	)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/policies/%d", ServerAddress, AgentID),
		httpmock.NewStringResponder(
			200,
			`[
				{"mobile_number": "1234567890", "premium": 2000, "type": "homeowner"},
				{"mobile_number": "123 456 7891", "premium": 200, "type": "renter"},
				{"mobile_number": "123-456-7892", "premium": 1500, "type": "homeowner"},
				{"mobile_number": "(123) 456-7893", "premium": 155, "type": "personal_auto"},
				{"mobile_number": "123-456-7894", "premium": 1000, "type": "homeowner"},
				{"mobile_number": "123-456-7890", "premium": 500, "type": "personal_auto"},
				{"mobile_number": "1234567892", "premium": 100, "type": "life"},
				{"mobile_number": "(123)456-7892", "premium": 200, "type": "homeowner"}
			]`,
		),
	)
}

package version

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	utils "github.com/nmccready/lambda-go-samples/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{HTTPMethod: "GET", Path: "/version"},
			expect:  utils.GetVersionMut(&utils.Version),
			err:     nil,
		},
	}

	for _, test := range tests {
		response, err := Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}

}

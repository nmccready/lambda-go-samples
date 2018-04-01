package main

import (
	"errors"
	"log"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

const DEFAULT_RESPONSE string = "done"

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("[v0.4] Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	log.Printf("[v0.4] BODY: %s\n", request.Body)

	m, err := url.ParseQuery(request.Body)
	if err != nil {
		panic(err)
	}
	resp := m["response_url"][0]
	command := m["command"][0]
	text := m["text"][0]
	log.Println("response_url", resp, "command:", command, "text:", text)

	awsInsatncesMsg(resp, strings.Contains(text, "ascii"))
	return events.APIGatewayProxyResponse{
		Body:       DEFAULT_RESPONSE,
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)

}

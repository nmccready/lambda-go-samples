package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided   = errors.New("no name was provided in the HTTP body")
	ErrInvalidGetRequest = errors.New("invali GET request")
)

const version = "0.0.1"

func getVersionJson() string {
	return fmt.Sprintf("{ \"version\": \"%v\" }", version)
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// log.Println("ENV", os.Environ())

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	log.Println("method:", request.HTTPMethod, "path:", request.Path, "res:", request.Resource)

	if request.HTTPMethod == "GET" {
		if request.Path == "version" {

			return events.APIGatewayProxyResponse{
				Body:       getVersionJson(),
				StatusCode: 200,
			}, nil
		}
		return events.APIGatewayProxyResponse{}, ErrInvalidGetRequest
	}

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)
}

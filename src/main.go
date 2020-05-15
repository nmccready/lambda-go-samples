package main

import (
	"errors"
	"log"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/nmccready/lambda-go-samples/src/utils"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided   = errors.New("no name was provided in the HTTP body")
	ErrInvalidGetRequest = errors.New("invalid GET request")
	Version              string
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// log.Println("ENV", os.Environ())

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	log.Println("method:", request.HTTPMethod, "path:", request.Path, "res:", request.Resource)

	if request.HTTPMethod == "GET" {
		if request.Path == "/version" {

			return events.APIGatewayProxyResponse{
				Body:       utils.GetVersionMut(&Version),
				StatusCode: 200,
			}, nil
		}

		if request.Path == "/tmp" {
			cmd := exec.Command("ls", "-la", "/tmp")
			outBytes, err := cmd.Output()

			if err != nil {
				return events.APIGatewayProxyResponse{StatusCode: 500}, err
			}

			out := string(outBytes)
			log.Println(out)

			return events.APIGatewayProxyResponse{
				Body:       out,
				StatusCode: 200,
			}, nil

		}
		return events.APIGatewayProxyResponse{StatusCode: 404}, ErrInvalidGetRequest
	}

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{StatusCode: 500}, ErrNameNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)
}

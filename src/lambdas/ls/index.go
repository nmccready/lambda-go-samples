package main

import (
	"log"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/nmccready/lambda-go-samples/src/utils"
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	utils.LogRequest(request)

	if request.Path != "/tmp" {
		return events.APIGatewayProxyResponse{StatusCode: 404}, utils.ErrInvalidGetRequest
	}

	dir := "/tmp"
	if request.QueryStringParameters["dir"] != "" {
		dir = request.QueryStringParameters["dir"]
	}

	cmd := exec.Command("ls", "-la", dir)
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

func main() {
	lambda.Start(Handler)
}

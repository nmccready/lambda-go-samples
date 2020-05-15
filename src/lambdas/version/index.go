package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/nmccready/lambda-go-samples/src/utils"
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	utils.LogRequest(request)

	log.Println("utils.Version ", utils.Version)

	if request.Path != "/version" {
		return events.APIGatewayProxyResponse{StatusCode: 404}, utils.ErrInvalidGetRequest
	}
	return events.APIGatewayProxyResponse{
		Body:       utils.GetVersionMut(&utils.Version),
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)
}

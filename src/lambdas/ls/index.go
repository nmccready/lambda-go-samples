package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/nmccready/lambda-go-samples/src/utils"
)

// Endpoints intent is to have fun and see the security limits of the filesystem of AWS lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	utils.LogRequest(request)

	if request.Path != "/ls" {
		return events.APIGatewayProxyResponse{StatusCode: 404}, utils.ErrInvalidGetRequest
	}
	out, err := utils.Ls(request.QueryStringParameters["dir"])
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	log.Println(out)

	return events.APIGatewayProxyResponse{
		Body:       out,
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)
}

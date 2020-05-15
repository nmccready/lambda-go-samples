package utils

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func LogRequest(request events.APIGatewayProxyRequest) {
	// log.Println("ENV", os.Environ())

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	log.Println("method:", request.HTTPMethod, "path:", request.Path, "res:", request.Resource)
}

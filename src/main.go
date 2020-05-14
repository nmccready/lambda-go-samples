package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided   = errors.New("no name was provided in the HTTP body")
	ErrInvalidGetRequest = errors.New("invalid GET request")
)

type PackageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func __Dirname() string {
	_, thisFileName, _, _ := runtime.Caller(1)
	return path.Dir(thisFileName)
}

func getFileBytes(filename string) []byte {
	filepath := path.Join(path.Dir(__Dirname()), filename)

	jsonFile, err := os.Open(filepath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	return bytes
}

func getVersionJson() string {
	var pkg PackageJson
	err := json.Unmarshal(getFileBytes("package.json"), &pkg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("{ \"version\": \"%v\" }", pkg.Version)
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
		if request.Path == "/version" {

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

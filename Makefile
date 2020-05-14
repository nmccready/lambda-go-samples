.PHONY: help, version, zip, build-osx, update-fn-code, deploy-stack, print-api-url
# SEE .env FOR $AWS_STACK_NAME, $AWS_S3_BUCKET

VERSION=$(shell jq -r ".version" package.json)

help: ## Generates this help message
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

version: ## Echo the current version
	@echo $(VERSION)

zip: ## builds the linux binary, and creates the zip for lambda upload
	GOOS=linux go build -o dist/main  -ldflags '-X main.Version="$(VERSION)"' src/main.go
	# GOOS=linux go build -o main  -ldflags '-X main.Version="$(VERSION)"' main.go aws.go
	aws cloudformation package --template-file template.yml --s3-bucket $(AWS_S3_BUCKET) --output-template-file packaged.yml

build: ## Build binary osx
	go build -o dist/main  -ldflags '-X main.Version="$(VERSION)"' src/main.go
	# go build -o main  -ldflags '-X main.Version="$(VERSION)"' main.go aws.go

update: zip ## Updates the lambda code in the existing CF stack
	aws lambda update-function-code --function-name $(shell aws cloudformation list-stack-resources --stack-name $(AWS_STACK_NAME) --query 'StackResourceSummaries[?ResourceType == `AWS::Lambda::Function`].PhysicalResourceId' --out text) \
	  --s3-bucket $(AWS_S3_BUCKET) \
	  --s3-key $(shell sed -n '/CodeUri/ s:.*/::p'  packaged.yml)

deploy: zip ## Deploys/Updates the cloudformation stack 
	aws cloudformation deploy \
	    --stack-name $(AWS_STACK_NAME) \
	    --template-file ./packaged.yml  \
	    --capabilities CAPABILITY_IAM

undeploy:
	aws cloudformation delete-stack --stack-name $(AWS_STACK_NAME)

get-api: ## Prints ApiGateway url base
	@aws cloudformation describe-stacks \
	    --stack-name $(AWS_STACK_NAME) \
	    --query 'Stacks[0].Outputs[0].OutputValue' \
	    --out text

test: ## test the module
	@npm test


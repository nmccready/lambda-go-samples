#AWS_STACK_NAME=lalyos-lambda-go
#AWS_S3_BUCKET=lp-lambda-go
VERSION = 0.0.6

zip: build
	aws cloudformation package --template-file template.yml --s3-bucket $(AWS_S3_BUCKET) --output-template-file packaged.yml

build:
	GOOS=linux go build -o main  -ldflags '-X main.Version="$(VERSION)"' main.go aws.go

build-osx:
	go build -o main  -ldflags '-X main.Version="$(VERSION)"' main.go aws.go

update: zip
	aws lambda update-function-code --function-name $(shell aws cloudformation list-stack-resources --stack-name $(AWS_STACK_NAME) --query 'StackResourceSummaries[?ResourceType == `AWS::Lambda::Function`].PhysicalResourceId' --out text) \
	  --s3-bucket $(AWS_S3_BUCKET) \
	  --s3-key $(shell sed -n '/CodeUri/ s:.*/::p'  packaged.yml)

deploy-stack: update
	aws cloudformation deploy \
	    --stack-name $(AWS_STACK_NAME) \
	    --template-file ./packaged.yml  \
	    --capabilities CAPABILITY_IAM

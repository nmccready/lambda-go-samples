## lambda-go-samples

An example of using AWS Lambda with Go

## Compile

```
GOOS=linux go build -o main
```

## Deployment

```
AWS_STACK_NAME=lalyos-lambda-go

aws cloudformation package --template-file template.yml --s3-bucket lp-lambda-go --output-template-file packaged.yml
aws cloudformation deploy --template-file /Users/lalyos/go/src/github.com/lalyos/lambdatest/packaged.yml --stack-name $AWS_STACK_NAME```
```

## Invoke API

```
AWS_API_ID=$(aws apigateway get-rest-apis --query 'items[? name == `'$AWS_STACK_NAME'`].id' --out text)
AWS_STAGE=$(aws apigateway get-stages --rest-api-id $AWS_API_ID --query 'item[0].stageName' --out text)

curl https://$AWS_API_ID.execute-api.$AWS_DEFAULT_REGION.amazonaws.com/$AWS_STAGE/ -d geza
```

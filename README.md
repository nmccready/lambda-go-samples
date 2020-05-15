# lambda-go-samples

see: `make help`

## versioning

[npm is gospel](https://docs.npmjs.com/cli/version)! **No code is better than code.** `npm version` command is more than sufficient and there is no need for any other sever code or logic

## dependencies

- `npm install -g foreman` to inject env variables like S3 buckets etc.
- `brew install jq`

## deployment example

`nf run make deploy`

## un-deploy (destroy stack)

`nf run make undeploy`

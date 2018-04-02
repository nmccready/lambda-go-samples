init() {
    cmd-export-ns aws "aws namespace"
    cmd-export aws-logstreams
    cmd-export aws-logs
    cmd-export aws-listStacks
    
    #AWS=$(which aws)
    deps-require aws
    AWS=.gun/bin/aws
}
aws-loggroups() {
  declare desc="lists aws loggroups"
  declare groupPattern=${1}

  $AWS logs describe-log-groups \
     --query 'logGroups[?contains(logGroupName,`'${groupPattern}'`)].logGroupName' \
     --out text
}

aws-logstreams() {
  declare desc="lists log streams belonging to a log group pattern"
  declare groupPattern=${1:? groupPattern required}

  groupName=$(aws-loggroups $groupPattern)
  $AWS logs describe-log-streams \
      --log-group-name $groupName \
      --query 'logStreams[-1].logStreamName' \
      --out text
}

aws-logs() {
  declare desc="list log messages belonging to a log group pattern"
  declare groupPattern=${1:? groupPattern required}
  declare messagePattern=${2}

  groupName=$(aws-loggroups $groupPattern)
  streamName=$(aws-logstreams $groupName)
  $AWS logs filter-log-events \
    --log-group-name $groupName \
    --log-stream-names $streamName \
    --query 'events[?contains(message,`'$messagePattern'`)].message' \
    --out text
}

aws-listStacks() {
  declare desc="List cloudformation stacks with CREATE_XXX state"

  $AWS cloudformation list-stacks \
    --query 'StackSummaries[?contains(StackStatus,`CREATE`)].StackName'
}
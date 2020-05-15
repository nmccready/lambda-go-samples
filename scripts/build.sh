#!/usr/bin/env bash

set -e -o pipefail

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$MY_DIR"
cd ../ # normalize to a known working directory could be ../../

VERSION=$(jq -r ".version" package.json)

# echo VERSION: $VERSION

LD_FLAGS="-X github.com/nmccready/lambda-go-samples/src/utils.Version=$VERSION"

# echo LD_FLAGS: $LD_FLAGS

for i in `ls -1 ./src/lambdas/**/*.go | grep -v test`; do
  binFileName=$(basename `dirname $i`)
  go build -o bin/$binFileName  -ldflags "$LD_FLAGS" $i
done

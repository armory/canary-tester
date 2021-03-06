#!/bin/bash -ex
cd "$(dirname "$0")/.."

rm -dfr build/
mkdir -p build/

docker run -v "$(pwd):/app" -w="/app" golang:1.9 \
  bash -c "go get github.com/DataDog/datadog-go/statsd && go build -o build/canary-tester cmd/canary-tester.go"

GIT_HASH=$(git rev-parse HEAD | cut -c -7)
docker run --rm \
  -e "BUILD_NUMBER=${BUILD_NUMBER}" \
  -e "BRANCH_NAME=$(echo ${GIT_HASH} | tr -d -- -)" \
  -v $(pwd):/app \
  -w="/app" \
  frekele/gradle:2.12-jdk8 \
  gradle -b /app/build.gradle buildDeb

echo "Built:"
ls -ahl build/distributions/
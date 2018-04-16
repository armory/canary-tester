#!/bin/bash -e
cd "$(dirname "$0")/.."

export GIT_HASH=$(git rev-parse HEAD | cut -c -7)
docker build -t armory/canary-tester:${GIT_HASH} .

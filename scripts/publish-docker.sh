#!/bin/bash -e
cd "$(dirname "$0")/.."

export GIT_HASH=$(git rev-parse HEAD | cut -c -7)
docker push armory/canary-tester:${GIT_HASH}
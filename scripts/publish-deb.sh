#!/bin/bash

VERSION=0.${CI_BUILD_NUMBER}.0-h${CI_BUILD_NUMBER}.${GIT_HASH}
DEB_FILE=canary-test_${VERSION}_all.deb
METADATA="deb_distribution=trusty;deb_component=main;publish=1;deb_architecture=i386,amd64"

REPO="debians-dev"
if [[ "${BRANCH_NAME}" == "master" ]]; then
  REPO="debians"
fi

BINTRAY_URL=https://api.bintray.com/content/armory/${REPO}

curl -T ./build/distributions/${DEB_FILE} \
  -u${BINTRAY_USER}:${BINTRAY_APIKEY} \
   "${BINTRAY_URL}/${APP_NAME}/${VERSION}/${DEB_FILE};${METADATA}"
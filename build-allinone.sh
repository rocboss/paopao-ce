#!/bin/sh
# eg.1 : sh build-image.sh
# eg.2, set image: sh build-image.sh bitbus/paopao-ce

VERSION=`git describe --tags --always | cut -f1,2 -d "-"` # eg.: 0.2.5
IMAGE="bitbus/paopao-ce"

if [ -n "$1" ]; then
  IMAGE="$1"
fi
if [ -n "$2" ]; then
  VERSION="$2"
fi

# build image
docker buildx build \
  --build-arg USE_DIST="yes" \
  --tag "$IMAGE:all-in-one-${VERSION}" \
  --tag "$IMAGE:all-in-one-latest" \
  . -f Dockerfile.allinone

# push to image rep
# if [ -n "$1" ]; then
#   docker push "$IMAGE:${VERSION}"
#   docker push "$IMAGE:latest"
# fi

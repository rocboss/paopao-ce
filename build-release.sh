#!/bin/sh
# eg.1 : sh build-image.sh
# eg.2, set tags: sh build-image.sh 'embed go_json'

TAGS='embed go_json'

if [ -n "$1" ]; then
  TAGS="$1"
fi

make release CGO_ENABLED=0 TAGS="$TAGS"

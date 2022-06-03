#!/usr/bin/env bash
# shellcheck disable=SC2155
set -e
DIST_PREFIX=${1}
DEBUG_MODE=${2}
TARGET_DIR="dist"
PLATFORMS="darwin/amd64 darwin/arm64 linux/amd64 windows/amd64"

rm -rf ${TARGET_DIR}
mkdir ${TARGET_DIR}

for pl in ${PLATFORMS}; do
    export CGO_ENABLED=0
    export GOOS=$(echo "${pl}" | cut -d'/' -f1)
    export GOARCH=$(echo "${pl}" | cut -d'/' -f2)
    export TARGET=${TARGET_DIR}/${DIST_PREFIX}_${GOOS}_${GOARCH}
    if [ "${GOOS}" == "windows" ]; then
        export TARGET=${TARGET_DIR}/${DIST_PREFIX}_${GOOS}_${GOARCH}.exe
    fi

    echo "build => ${TARGET}"
    if [ "${DEBUG_MODE}" == "debug" ]; then
        go build -trimpath -gcflags "all=-N -l" -o "${TARGET}" -tags "${TAGS}" \
            -ldflags "-X 'main.version=${BUILD_VERSION}' \
                        -X 'main.buildDate=${BUILD_DATE}' \
                        -X 'main.commitID=${SHA_SHORT}'\
                        -w -s"
    else
        go build -trimpath -o "${TARGET}" -tags "${TAGS}" \
            -ldflags "-X 'main.version=${BUILD_VERSION}' \
                        -X 'main.buildDate=${BUILD_DATE}' \
                        -X 'main.commitID=${SHA_SHORT}'\
                        -w -s"
    fi
done

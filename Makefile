.PHONY: all build run clean fmt help
BUILD_VERSION   := $(shell cat version)
BUILD_DATE := $(shell date +'%Y-%m-%d %H:%M:%S')
SHA_SHORT := $(shell git rev-parse --short HEAD)
all: fmt build
build:
	@go mod download
	@echo Build paopao-ce
	bash build.sh paopao-ce
run:
	@go run -ldflags "-X 'main.version=${BUILD_VERSION}' -X 'main.buildDate=${BUILD_DATE}' -X 'main.commitID=${SHA_SHORT}'" .
clean:
	@go clean
	@find ./dist -type f -exec rm -r {} +
	@find ./tmp -type f -exec rm -r {} +
fmt:
	@echo Formatting...
	@go fmt ./global/...
	@go fmt ./internal/...
	@go fmt ./pkg/...
	@go vet -composites=false ./global/...
	@go vet -composites=false ./internal/...
	@go vet -composites=false ./pkg/...
help:
	@echo "make: make"
	@echo "make build: build executables"
.EXPORT_ALL_VARIABLES:
GO111MODULE = on

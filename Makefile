.PHONY: all build build-web run test clean fmt pre-commit help

PROJECT = paopao-ce
TARGET = paopao
ifeq ($(OS),Windows_NT)
TARGET := $(TARGET).exe
endif
TARGET_BIN = $(basename $(TARGET))

ifeq (n$(CGO_ENABLED),n)
CGO_ENABLED := 1
endif

RELEASE_ROOT = release
RELEASE_FILES = LICENSE README.md CHANGELOG.md config.yaml.sample docker-compose.yaml scripts docs
RELEASE_LINUX_AMD64 = $(RELEASE_ROOT)/linux-amd64/$(PROJECT)
RELEASE_DARWIN_AMD64 = $(RELEASE_ROOT)/darwin-amd64/$(PROJECT)
RELEASE_DARWIN_ARM64 = $(RELEASE_ROOT)/darwin-arm64/$(PROJECT)
RELEASE_WINDOWS_AMD64 = $(RELEASE_ROOT)/windows-amd64/$(PROJECT)

BUILD_VERSION := $(shell git describe --tags --always | cut -f1 -f2 -d "-")
BUILD_DATE := $(shell date +'%Y-%m-%d %H:%M:%S')
SHA_SHORT := $(shell git rev-parse --short HEAD)

TAGS = ""
MOD_NAME = github.com/rocboss/paopao-ce
LDFLAGS = -X "${MOD_NAME}/pkg/version.version=${BUILD_VERSION}" \
          -X "${MOD_NAME}/pkg/version.buildDate=${BUILD_DATE}" \
          -X "${MOD_NAME}/pkg/version.commitID=${SHA_SHORT}" -w -s

all: fmt build

build:
	@go mod download
	@echo Build paopao-ce
	@go build -pgo=auto -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(RELEASE_ROOT)/$(TARGET)

build-web:
	@cd web && rm -rf dist/* && yarn build && cd -

run:
	@go run -pgo=auto -trimpath -gcflags "all=-N -l" -tags '$(TAGS)' -ldflags '$(LDFLAGS)' . serve

.PHONY: release
release: linux-amd64 darwin-amd64 darwin-arm64 windows-x64
	@echo Package paopao-ce
	@cp -rf $(RELEASE_FILES) $(RELEASE_LINUX_AMD64)
	@cp -rf $(RELEASE_FILES) $(RELEASE_DARWIN_AMD64)
	@cp -rf $(RELEASE_FILES) $(RELEASE_DARWIN_ARM64)
	@cp -rf $(RELEASE_FILES) $(RELEASE_WINDOWS_AMD64)
	@cd $(RELEASE_LINUX_AMD64)/.. && rm -f *.zip && zip -r $(PROJECT)-linux_amd64.zip $(PROJECT) && cd -
	@cd $(RELEASE_DARWIN_AMD64)/.. && rm -f *.zip && zip -r $(PROJECT)-darwin_amd64.zip $(PROJECT) && cd -
	@cd $(RELEASE_DARWIN_ARM64)/.. && rm -f *.zip && zip -r $(PROJECT)-darwin_arm64.zip $(PROJECT) && cd -
	@cd $(RELEASE_WINDOWS_AMD64)/.. && rm -f *.zip && zip -r $(PROJECT)-windows_amd64.zip $(PROJECT) && cd -

.PHONY: linux-amd64
linux-amd64:
	@echo Build paopao-ce [linux-amd64] CGO_ENABLED=$(CGO_ENABLED) TAGS="'$(TAGS)'"
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -pgo=auto -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(RELEASE_LINUX_AMD64)/$(TARGET_BIN)

.PHONY: darwin-amd64
darwin-amd64:
	@echo Build paopao-ce [darwin-amd64] CGO_ENABLED=$(CGO_ENABLED) TAGS="'$(TAGS)'"
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=darwin GOARCH=amd64 go build -pgo=auto -trimpath  -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(RELEASE_DARWIN_AMD64)/$(TARGET_BIN)

.PHONY: darwin-arm64
darwin-arm64:
	@echo Build paopao-ce [darwin-arm64] CGO_ENABLED=$(CGO_ENABLED) TAGS="'$(TAGS)'"
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=darwin GOARCH=arm64 go build -pgo=auto -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(RELEASE_DARWIN_ARM64)/$(TARGET_BIN)

.PHONY: windows-x64
windows-x64:
	@echo Build paopao-ce [windows-x64] CGO_ENABLED=$(CGO_ENABLED) TAGS="'$(TAGS)'"
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=windows GOARCH=amd64 go build -pgo=auto -trimpath  -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(RELEASE_WINDOWS_AMD64)/$(TARGET_BIN).exe

.PHONY: generate
generate: gen-mir gen-rpc gen-sqlc gen-yesql

.PHONY: gen-mir
gen-mir:
	@go generate mirc/gen.go
	@go fmt ./auto/api/...

.PHONY: gen-yesql
gen-yesql:
	@go generate internal/dao/sakila/yesql/gen.go
	@go fmt ./internal/dao/sakila/yesql/...

.PHONY: gen-rpc
gen-rpc:
	@rm -rf auto/rpc auto/connect
	@buf generate proto
	@go fmt ./auto/rpc/... ./auto/connect/...

.PHONY: gen-sqlc
gen-sqlc:
	@find internal/dao/slonik/sqlc/postgres -name '*.go' -exec rm -f {} +
	@go generate internal/dao/slonik/sqlc/gen.go
	@go fmt ./internal/dao/slonik/sqlc/...

.PHONY: proto-mod
proto-mod:
	@cd proto/ && buf mod update

.PHONY: proto-lint
proto-lint:
	@cd proto/ && buf lint

clean:
	@go clean
	@find ./release -type f -exec rm -r {} +

fmt:
	@echo Formatting...
	@go fmt ./internal/...
	@go fmt ./pkg/...
	@go vet -composites=false ./internal/...
	@go vet -composites=false ./pkg/...

test:
	@go test ./...

pre-commit: fmt
	@go mod tidy


.PHONY: install-plugins
install-plugins: install-protobuf-plugins

.PHONY: install-protobuf-plugins
install-protobuf-plugins:
	@go install github.com/bufbuild/buf/cmd/buf@v1.25.0
	@go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@v1.25.0
	@go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@v1.25.0
	@go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
	@go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: install-sqlc-plugins
install-sqlc-plugins:
	@go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

help:
	@echo "make: make"
	@echo "make run: start api server"
	@echo "make build: build executable"
	@echo "make release: build release executables"
	@echo "make run TAGS='embed': start api server and serve embed web frontend"
	@echo "make build TAGS='embed': build executable with embed web frontend"
	@echo "make release TAGS='embed': build release executables with embed web frontend"

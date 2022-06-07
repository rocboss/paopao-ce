.PHONY: all build run test clean fmt help

TARGET = paopao-ce
RELEASE_ROOT = release
RELEASE_FILES = LICENSE README.md config.yaml.sample scripts assets configs
RELEASE_LINUX_AMD64 = $(RELEASE_ROOT)/linux-amd64/$(TARGET)
RELEASE_DARWIN_AMD64 = $(RELEASE_ROOT)/darwin-amd64/$(TARGET)
RELEASE_DARWIN_ARM64 = $(RELEASE_ROOT)/darwin-arm64/$(TARGET)
RELEASE_WINDOWS_AMD64 = $(RELEASE_ROOT)/windows-amd64/$(TARGET)

BUILD_VERSION   := $(shell cat version)
BUILD_DATE := $(shell date +'%Y-%m-%d %H:%M:%S')
SHA_SHORT := $(shell git rev-parse --short HEAD)

TAGS = ""
LDFLAGS = -X "main.version=${BUILD_VERSION}" -X "main.buildDate=${BUILD_DATE}" -X "main.commitID=${SHA_SHORT}" -w -s

all: fmt build

build:
	@go mod download
	@echo Build paopao-ce
	@go build -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(RELEASE_ROOT)/$(TARGET)

run:
	@go run -trimpath -gcflags "all=-N -l" -tags '$(TAGS)' -ldflags '$(LDFLAGS)' .

.PHONY: release
release: linux-amd64 darwin-amd64 darwin-arm64 windows-x64
	@echo Package paopao-ce
	@cp -rf $(RELEASE_FILES) $(RELEASE_LINUX_AMD64)
	@cp -rf $(RELEASE_FILES) $(RELEASE_DARWIN_AMD64)
	@cp -rf $(RELEASE_FILES) $(RELEASE_DARWIN_ARM64)
	@cp -rf $(RELEASE_FILES) $(RELEASE_WINDOWS_AMD64)
	@cd $(RELEASE_LINUX_AMD64)/.. && rm -f *.zip && zip -r $(TARGET)-linux_amd64.zip $(TARGET) && cd -
	@cd $(RELEASE_DARWIN_AMD64)/.. && rm -f *.zip && zip -r $(TARGET)-darwin_amd64.zip $(TARGET) && cd -
	@cd $(RELEASE_DARWIN_ARM64)/.. && rm -f *.zip && zip -r $(TARGET)-darwin_arm64.zip $(TARGET) && cd -
	@cd $(RELEASE_WINDOWS_AMD64)/.. && rm -f *.zip && zip -r $(TARGET)-windows_amd64.zip $(TARGET) && cd -

.PHONY: linux-amd64
linux-amd64:
	@echo Build paopao-ce [linux-amd64]
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(RELEASE_LINUX_AMD64)/$(TARGET)

.PHONY: darwin-amd64
darwin-amd64:
	@echo Build paopao-ce [darwin-amd64]
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -trimpath  -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(RELEASE_DARWIN_AMD64)/$(TARGET)

.PHONY: darwin-arm64
darwin-arm64:
	@echo Build paopao-ce [darwin-arm64]
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(RELEASE_DARWIN_ARM64)/$(TARGET)

.PHONY: windows-x64
windows-x64:
	@echo Build paopao-ce [windows-x64]
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath  -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(RELEASE_WINDOWS_AMD64)/$(TARGET).exe

clean:
	@go clean
	@find ./release -type f -exec rm -r {} +

fmt:
	@echo Formatting...
	@go fmt ./global/...
	@go fmt ./internal/...
	@go fmt ./pkg/...
	@go vet -composites=false ./global/...
	@go vet -composites=false ./internal/...
	@go vet -composites=false ./pkg/...

test:
	@go test ./...

help:
	@echo "make: make"
	@echo "make run: start api server"
	@echo "make build: build executable with debug information"
	@echo "make release: build release executables"
	@echo "make run TAGS='embed': start api server and serve embed web frontend"
	@echo "make build TAGS='embed': build executable with embed web frontend"
	@echo "make release TAGS='embed': build release executables with embed web frontend"

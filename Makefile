ROOT_DIR := $(shell pwd)
SHELL := /bin/sh

GO_FILES = $(shell find $(ROOT_DIR) -name "*.go" ! -path "*/thrift_0_9_2/*")
GO_MODULES = $(shell go list ./... | grep -v "thrift_0_9_2" | grep -v "e2e")
COVERAGE_DIR ?= $(ROOT_DIR)/coverage
BUILDER ?= ./ocb
VERSION ?= 0.113.0

export GO111MODULE = on

default: all

all: clean format lint test coverage build

clean:
	rm -rf otel build ocb coverage

format:
	test -z "$(shell gofmt -l $(GO_FILES))" || (echo "Run 'make format-fix' to format code" && exit 1)

format-fix:
	gofmt -s -w $(GO_FILES)

lint:
	golangci-lint run

test:
	go test -race -v -failfast $(GO_MODULES)

coverage:
	mkdir -p $(COVERAGE_DIR)
	go test -coverpkg=./... -coverprofile=$(COVERAGE_DIR)/coverage.out $(GO_MODULES)
	go tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/index.html

build:
	CGO_ENABLED=0 go build $(BUILD_FLAGS) -ldflags "$(LDFLAGS)"

document:
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	gomarkdoc --output ./docs/receiver.md ./

download-builder:
	@if [ "$(BUILDER)" = "./ocb" ] && [ ! -x $(BUILDER) ]; then \
		sys=$$(uname -s | tr '[:upper:]' '[:lower:]'); \
		arch=$$(uname -m | sed 's/x86_64/amd64/'); \
		curl -o $(BUILDER) -L "https://github.com/open-telemetry/opentelemetry-collector-releases/releases/download/cmd/builder/v$(VERSION)/ocb_$(VERSION)_$${sys}_$${arch}"; \
		chmod +x $(BUILDER); \
	fi

build-local: download-builder
	mkdir -p collector
	CGO_ENABLED=0 $(BUILDER) --config config/build_config.yaml

run: build-local
	./collector --config config/run_config.yaml

debug: build-local
	dlv --listen=:2345 --api-version=2 --headless=true --accept-multiclient --log exec ./collector -- --config config/run_config.yaml

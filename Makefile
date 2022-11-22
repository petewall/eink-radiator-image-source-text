HAS_GINKGO := $(shell command -v ginkgo;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_COUNTERFEITER := $(shell command -v counterfeiter;)
PLATFORM := $(shell uname -s)

# #### DEPS ####
.PHONY: deps-counterfeiter deps-ginkgo deps-modules

deps-counterfeiter:
ifndef HAS_COUNTERFEITER
	go install github.com/maxbrunsfeld/counterfeiter/v6@latest
endif

deps-ginkgo:
ifndef HAS_GINKGO
	go install github.com/onsi/ginkgo/v2/ginkgo
endif

deps-modules:
	go mod download

# #### TEST ####
.PHONY: lint test

lint: deps-modules
ifndef HAS_GOLANGCI_LINT
ifeq ($(PLATFORM), Darwin)
	brew install golangci-lint
endif
ifeq ($(PLATFORM), Linux)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
endif
endif
	golangci-lint run

test: deps-modules deps-ginkgo
ifeq ($(PLATFORM), Linux)
	mkdir -p ~/.fonts
	cp fonts/ubuntu/*.ttf ~/.fonts
endif
	ginkgo -r .

TEST_CONFIGS = $(shell find test/inputs -name "*.yaml" -type f)
TEST_OUTPUTS = $(TEST_CONFIGS:test/inputs/%.yaml=test/outputs/%.png)
test-artifacts: build/text $(TEST_OUTPUTS)

test/outputs/%.png: test/inputs/%.yaml
	./build/text generate --config $? --output test/outputs/$(basename $(shell basename $@)).png --height 300 --width 400

# #### BUILD ####
.PHONY: build
SOURCES = $(shell find . -name "*.go" | grep -v "_test\." )
VERSION := $(or $(VERSION), dev)
LDFLAGS="-X github.com/petewall/eink-radiator-image-source-text/cmd.Version=$(VERSION)"

build: build/text

build/text: $(SOURCES) deps-modules
	go build -o $@ -ldflags ${LDFLAGS} github.com/petewall/eink-radiator-image-source-text

build-all: build/text-arm6 build/text-arm7 build/text-darwin-amd64

build/text-arm6: $(SOURCES) deps-modules
	GOOS=linux GOARCH=arm GOARM=6 go build -o $@ -ldflags ${LDFLAGS} github.com/petewall/eink-radiator-image-source-text

build/text-arm7: $(SOURCES) deps-modules
	GOOS=linux GOARCH=arm GOARM=7 go build -o $@ -ldflags ${LDFLAGS} github.com/petewall/eink-radiator-image-source-text

build/text-darwin-amd64: $(SOURCES) deps-modules
	GOOS=darwin GOARCH=amd64 go build -o $@ -ldflags ${LDFLAGS} github.com/petewall/eink-radiator-image-source-text

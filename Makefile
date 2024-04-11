# useful paths
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT_PATH := $(patsubst %/,%,$(dir $(MKFILE_PATH)))
PROJECT_BIN := $(PROJECT_PATH)/bin
PROJECT_DIST := $(PROJECT_PATH)/dist

OS_NAME := $(shell uname -s | tr A-Z a-z)
OS_ARCH := $(shell uname -m | tr A-Z a-z)

ifeq ($(OS_NAME),linux)
	OS_NAME = "linux"
endif
ifeq ($(OS_NAME),darwin)
	OS_NAME = "osx"
endif

ifeq ($(OS_ARCH),x86_64)
	OS_ARCH = x64
endif
ifneq ($(filter %86,$(OS_ARCH)),)
	OS_ARCH = x86
endif
ifneq ($(filter arm%,$(OS_ARCH)),)
	OS_ARCH = arm64
endif

# env variables
GOLANGCI_LINT ?= ${PROJECT_BIN}/golangci-lint
GOLANGCI_LINT_VERSION ?= "v1.54.2"
KIOTA ?= ${PROJECT_BIN}/kiota
KIOTA_VERSION ?= "v1.13.0"
HORREUM_BRANCH ?= "master"
HORREUM_OPENAPI_PATH ?= "https://raw.githubusercontent.com/Hyperfoil/Horreum/${HORREUM_BRANCH}/docs/site/content/en/openapi/openapi.yaml"
GENERATED_CLIENT_PATH = "${PROJECT_PATH}/pkg/raw_client"
OPENAPI_PATH = "${PROJECT_PATH}/openapi"
OPENAPI_SPEC = "${OPENAPI_PATH}/openapi.yaml"

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Cleanup

.PHONY: clean-bin
clean-bin: ## Clean external tools
	@rm -rf ${PROJECT_BIN}

.PHONY: clean
clean: ## Clean output directories
	@rm -rf ${PROJECT_DIST} ${GENERATED_CLIENT_PATH} ${OPENAPI_PATH}


##@ Development

.PHONY: kiota
kiota: ${PROJECT_BIN}/kiota ## Install kiota tool under ${PROJECT_PATH}/bin

${PROJECT_BIN}/kiota:
	@{\
		set -e ;\
		echo "installing kiota version ${KIOTA_VERSION}" ;\
		mkdir -p ${PROJECT_BIN}/kiota-installation ;\
		cd ${PROJECT_BIN}/kiota-installation ;\
		curl -sLO https://github.com/microsoft/kiota/releases/download/${KIOTA_VERSION}/${OS_NAME}-${OS_ARCH}.zip ;\
		unzip -o ${OS_NAME}-${OS_ARCH}.zip ;\
		mv kiota ${PROJECT_BIN}/ ;\
		rm -rf ${PROJECT_BIN}/kiota-installation ;\
	}

.PHONY: golangci-lint
golangci-lint: ${PROJECT_BIN}/golangci-lint ## Install golangci-lint tool under ${PROJECT_PATH}/bin

${PROJECT_BIN}/golangci-lint:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) ${GOLANGCI_LINT_VERSION}

.PHONY: tools
tools: kiota golangci-lint ## Install external tools

.PHONY: openapi
openapi: ${OPENAPI_PATH}/openapi.yaml ## Fetch openapi spec from Horreum

${OPENAPI_SPEC}:
	@if [ ! -f ${OPENAPI_SPEC} ]; then \
		mkdir -p ${OPENAPI_PATH} ;\
		echo "fetching openapi from ${HORREUM_OPENAPI_PATH}"; \
		curl -sSfL -o $@ ${HORREUM_OPENAPI_PATH}; \
	fi

.PHONY: generate-client
generate-client: tools ## Generate kiota Horreum client
	@{\
		set -e ;\
		${KIOTA} generate -l go -c HorreumRawClient -n github.com/hyperfoil/horreum-client-golang/pkg/raw_client -d ${OPENAPI_PATH}/openapi.yaml -o ${GENERATED_CLIENT_PATH} ;\
	}

.PHONY: generate
generate: ${OPENAPI_SPEC} generate-client ## Generate sources

.PHONY: test
test: generate ## Run unit tests, those without `integration` build tag
	go test ./... -v ${OTHER_PARAMS}

.PHONY: test-all
test-all: generate ## Run all tests, including integration ones
	go test ./... -tags=integration -v ${OTHER_PARAMS}

##@ Build

.PHONY: lint
lint: tools generate ## Lint pkg/horreum
	${GOLANGCI_LINT} run ./pkg/horreum/...

.PHONY: vet
vet: generate ## Examines source code and reports suspicious constructs
	go vet ./...

.PHONY: build
build: generate vet lint ## Build golang pkg
	go build ./pkg/...
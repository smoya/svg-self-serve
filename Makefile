# Makefile based on https://gist.github.com/thomaspoignant/5b72d579bd5f311904d973652180c705

GOCMD=go
PROJECT_NAME := $(shell basename "$(PWD)")
BINARY_NAME?=$(PROJECT_NAME)
BIN_DIR?=bin
FUNC_DIR?=functions
FUNC_OUTPUT_DIR?=$(FUNC_DIR)/bin

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build vendor

all: help

## Build:
build: ## Build your project and put the output binary in bin/out
	mkdir -p $(BIN_DIR)
	$(GOCMD) build -o bin/out/$(BINARY_NAME) .

build-functions: ## Build functions/lambdas your project and put the output binary in bin/out
	mkdir -p $(FUNC_OUTPUT_DIR)
	$(GOCMD) build -o $(FUNC_OUTPUT_DIR)/generate $(FUNC_DIR)/generate/main.go

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
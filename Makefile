SHELL := bash # we want bash behaviour in all shell invocations
PLATFORM := $(shell uname)
platform := $(shell echo $(PLATFORM) | tr A-Z a-z)

# https://stackoverflow.com/questions/4842424/list-of-ansi-color-escape-sequences
RED := \033[1;31m
GREEN := \033[1;32m
YELLOW := \033[1;33m
BOLD := \033[1m
NORMAL := \033[0m

HELP_TARGET_DEPTH ?= \#\#
.PHONY: help
help:
	@awk -F':+ |$(HELP_TARGET_DEPTH)' '/^[^.][0-9a-zA-Z._%-]+:+.+$(HELP_TARGET_DEPTH).+$$/ { printf "\033[36m%-26s\033[0m %s\n", $$1, $$3 }' $(MAKEFILE_LIST) \
	| sort

.DEFAULT_GOAL := help

### Tools and setup

GO ?= $(shell which go)

GINKGO ?= $(GOPATH)/bin/ginkgo
.PHONY: ginkgo
ginkgo: | $(GINKGO)

### build

build: ## Just compile all
	$(GO) build -v .

### Tests

GINKGO_RUN_SHARED_FLAGS := --randomize-all --race
GINKGO_RUN_FLAGS := -r -p

.PHONY: tests
tests: | $(GINKGO) ## Run unit tests. Make sure you install-tools before running this target. Use GINKGO_EXTRA to add extra flags to ginkgo
	@printf "$(GREEN)Running all tests in parallel$(NORMAL)\n"
	$(GINKGO) $(GINKGO_RUN_SHARED_FLAGS) $(GINKGO_RUN_FLAGS) --skip-package "integration" $(GINKGO_EXTRA) .

.PHONY: integration
integration: | $(GINKGO) ## Run all tests including integration
	@printf "$(GREEN)Running integration tests$(NORMAL)\n"
	$(GINKGO) $(GINKGO_RUN_SHARED_FLAGS) $(GINKGO_RUN_FLAGS) $(GINKGO_EXTRA) .
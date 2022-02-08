# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /usr/bin/env bash

export WORKDIR=$(shell pwd)
export APP_NAME=redis-viewer

.PHONY: build
# build executable file for dev
build:
	sh scripts/build.sh

.PHONY: run
# run executable file
run:
	sh output/run.sh

.PHONY: clean
# clean build cache and docker images
clean:
	sh scripts/clean.sh

.PHONY: generate
# run go generate
generate:
	go generate ./...

# show help
help:
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\w0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
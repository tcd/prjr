.DEFAULT_GOAL := help

GO ?= go
SHELL := /bin/sh
GOBIN_DIR=${GOBIN}
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

# =============================================================================
# Go
# =============================================================================

go-test:
	$(GO) test -v ./...

go-build:
	GO111MODULE=on $(GO) build -o build/prjr

go-clean:
	$(GO) clean ./...

# Install prjr to $GOBIN.
go-install:
	GO111MODULE=on $(GO) install

# Remove prjr from $GOBIN.
go-uninstall:
	@rm -f $(GOBIN_DIR)/$(PROJECT_NAME)

# =============================================================================
# Docker
# =============================================================================

docker-build:
	docker build --rm -t prjr-alpine .

docker-run:
	docker run -it --rm prjr-alpine

docker: docker-build docker-run

# =============================================================================
# Everything Else
# =============================================================================

build: go-build

clean: go-clean

test: go-test

install: go-install

uninstall: go-uninstall

.PHONY: clean test install uninstall
.PHONY: go-build go-clean go-test go-install go-uninstall
.PHONY: docker docker-build docker-run

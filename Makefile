.DEFAULT_GOAL := help

SHELL := /bin/sh
GOBIN_DIR=${GOBIN}
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

# =============================================================================
# Go
# =============================================================================

go-test:
	go test -v ./...

go-clean:
	go clean ./...

# Install prjr to $GOBIN.
go-install:
	go install

# Remove prjr from $GOBIN.
go-uninstall:
	@rm -f $(GOBIN_DIR)/$(PROJECT_NAME)

# Run cmd/prjr/*.go. 
# You'll be prompted for input to pass to the program.
cmd:	
	@scripts/cmd.sh

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

clean: go-clean

test: go-test

install: go-install

uninstall: go-uninstall

.PHONY: clean test install uninstall
.PHONY: cmd go-test go-clean go-install go-uninstall
.PHONY: docker docker-build docker-run

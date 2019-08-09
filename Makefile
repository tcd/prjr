.DEFAULT_GOAL := help

SHELL := /bin/bash
GOBIN_DIR=${GOBIN}
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

go-test:
	go test -v ./...

go-clean:
	go clean ./...

go-install:
	cd cmd/$(PROJECT_NAME) && go install

go-uninstall:
	@rm -f $(GOBIN_DIR)/$(PROJECT_NAME)

# Run cmd/prjr/*.go. 
# You'll be prompted for input to pass to the program.
run:	
	@scripts/cmd.sh

clean: go-clean

test: go-test

install: go-install

uninstall: go-uninstall

help:	
	@echo
	@echo "  test  – run 'go test' for the entire project"
	@echo "  clean – clean all files built by 'go build'"
	@echo "  run   – run cmd/$(PROJECT_NAME)/*.go"
	@echo

.PHONY: clean test help cmd
.PHONY: go-test go-clean go-install go-uninstall

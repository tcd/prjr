.DEFAULT_GOAL := help

SHELL := /bin/bash
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

go-test:
	go test -v ./...

go-clean:
	go clean ./...

# Run cmd/prjr/*.go. 
# You'll be prompted for commands to pass to the program.
cmd:	
	@scripts/cmd.sh

clean: go-clean

test: go-test

help:	
	@echo
	@echo "  test  – run 'go test' for the entire project"
	@echo "  clean – clean all files built by 'go build'"
	@echo

.PHONY: clean test help cmd go-test go-clean

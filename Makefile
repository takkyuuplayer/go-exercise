.PHONY: test run run-test

setup:
	go mod download

test:
	go test ./...

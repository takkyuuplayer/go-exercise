.PHONY: test

setup:
	go mod download

test:
	go clean -testcache
	go test ./...

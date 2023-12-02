GOIMPORTS:=go run golang.org/x/tools/cmd/goimports@latest
STATICCHECK:=go run honnef.co/go/tools/cmd/staticcheck@latest

SHELL=/bin/bash
.PHONY: test

compose/up:
	docker compose up -d --pull always --build

compose/down:
	docker compose down --remove-orphans

fmt:
	$(GOIMPORTS) -w .

staticcheck:
	$(STATICCHECK) ./...

update:
	go mod edit -go=$(shell go env GOVERSION | sed 's/^go//' | sed -e 's/.[0-9]$+$$//g')
	go get -u
	go mod tidy
	go mod download all

test:
	go clean -testcache
	set -a; source .env; go test ./...

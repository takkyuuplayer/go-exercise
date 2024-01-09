SHELL:=/bin/bash
GOIMPORTS:=go run golang.org/x/tools/cmd/goimports@latest
STATICCHECK:=go run honnef.co/go/tools/cmd/staticcheck@latest
TPARAGEN:=go run github.com/sho-hata/tparagen/cmd/tparagen@latest

compose/up:
	docker compose up -d --pull always --build

compose/down:
	docker compose down --remove-orphans

fmt:
	$(GOIMPORTS) -w .

staticcheck:
	$(STATICCHECK) ./...

tparagen:
	$(TPARAGEN)

update:
	go mod edit -go=$(shell go env GOVERSION | sed 's/^go//' | sed -e 's/.[0-9]$+$$//g')
	go get -u
	go mod tidy
	go mod download all

.PHONY: test
test:
	go clean -testcache
	set -a; source .env; go test ./...

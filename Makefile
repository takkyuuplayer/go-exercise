SHELL:=/bin/bash
GOIMPORTS:=go run golang.org/x/tools/cmd/goimports@latest
GOLANGCI_LINT:=go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
TPARAGEN:=go run github.com/sho-hata/tparagen/cmd/tparagen@latest

compose/up:
	docker compose up -d --pull always --build

compose/down:
	docker compose down --remove-orphans

fmt:
	$(GOIMPORTS) -w .

golangci-lint:
	$(GOLANGCI_LINT) run ./...

tparagen:
	$(TPARAGEN)

update:
	go get -u ./...
	go mod tidy
	go mod download

.PHONY: test
test:
	go clean -testcache
	set -a; source .env; go test ./...

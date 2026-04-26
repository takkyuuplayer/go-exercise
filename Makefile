SHELL:=/bin/bash
GOIMPORTS:=go run golang.org/x/tools/cmd/goimports@latest
GOLANGCI_LINT:=go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
TPARAGEN:=go run github.com/sho-hata/tparagen/cmd/tparagen@latest

compose/up:
	bin/find-free-ports.sh > .env
	docker compose up -d --pull always --build
	$(MAKE) .env

.PHONY: .env
.env:
	@set -a; source .env; set +a; \
	MYSQL_PORT=$$(docker compose port mysql 3306 | cut -d: -f2) \
	BIGQUERY_PORT=$$(docker compose port bigquery 9050 | cut -d: -f2) \
	envsubst < .env.template > .env

compose/down:
	docker compose down --remove-orphans

fmt:
	$(GOIMPORTS) -w .

golangci-lint:
	$(GOLANGCI_LINT) run ./...

tparagen:
	$(TPARAGEN)

update:
	go mod edit -go=$(shell go env GOVERSION | sed 's/^go//' | sed -e 's/.[0-9]$+$$//g')
	go get -u
	go mod tidy
	go mod download

load/etc/hosts/redis-cluster:
	@echo "127.0.0.1 redis-cluster-node-1 redis-cluster-node-2 redis-cluster-node-3" | sudo tee -a /etc/hosts

unload/etc/hosts/redis-cluster:
	@sudo sed -i '' '/^127\.0\.0\.1 redis-cluster-node-1 redis-cluster-node-2 redis-cluster-node-3$$/d' /etc/hosts

.PHONY: test
test:
	go clean -testcache
	set -a; source .env; go test ./...

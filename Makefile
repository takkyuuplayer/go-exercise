SHELL:=/bin/bash

# compose のプロジェクト名はディレクトリの basename 由来のため、git worktree を
# 複数切ると衝突しうる。パスのハッシュを混ぜて worktree ごとに一意にする。
COMPOSE_PROJECT_NAME:=$(shell basename '$(CURDIR)' | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9]/-/g')-$(shell printf %s '$(CURDIR)' | shasum | cut -c1-8)
export COMPOSE_PROJECT_NAME

GOIMPORTS:=go run golang.org/x/tools/cmd/goimports@latest
GOLANGCI_LINT:=go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
TPARAGEN:=go run github.com/sho-hata/tparagen/cmd/tparagen@latest

# .env は 2 回生成する:
#  up 前 … redis のポート変数を .env に埋めて compose の変数補間に渡すため
#  up 後 … docker が割り当てた mysql/bigquery のホストポートを .env に反映するため
compose/up:
	bin/find-free-ports.sh > .env.ports
	$(MAKE) .env
	docker compose up -d --pull always --build
	$(MAKE) .env

.PHONY: .env
.env:
	@set -a; source .env.ports; set +a; \
	MYSQL_PORT=$$(docker compose port mysql 3306 2>/dev/null | cut -d: -f2) \
	BIGQUERY_PORT=$$(docker compose port bigquery 9050 2>/dev/null | cut -d: -f2) \
	envsubst < .env.template > .env

compose/down:
	docker compose down --remove-orphans
	@rm -f .env .env.ports

fmt:
	$(GOIMPORTS) -w .

golangci-lint:
	$(GOLANGCI_LINT) run ./...

tparagen:
	$(TPARAGEN)

load/etc/hosts/redis-cluster:
	@echo "127.0.0.1 redis-cluster-node-1 redis-cluster-node-2 redis-cluster-node-3" | sudo tee -a /etc/hosts

unload/etc/hosts/redis-cluster:
	@sudo sed -i '' '/^127\.0\.0\.1 redis-cluster-node-1 redis-cluster-node-2 redis-cluster-node-3$$/d' /etc/hosts

.PHONY: test
test:
	go clean -testcache
	set -a; source .env; go test ./...

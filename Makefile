.PHONY: test

compose/up:
	docker compose up -d --pull --build

compose/down:
	docker compose down --remove-orphans

update:
	go mod edit -go=$(shell go env GOVERSION | sed 's/^go//' | sed -e 's/.[0-9]$+$$//g')
	go get -u all
	go mod tidy
	go mod download all

test:
	go clean -testcache
	go test ./...

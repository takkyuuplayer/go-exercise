.PHONY: test

compose/up:
	docker compose up -d --pull --build

compose/down:
	docker compose down --remove-orphans

test:
	go clean -testcache
	go test ./...

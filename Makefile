.PHONY: test run run-test

setup:
	go mod download

test:
	go test ./...

run:
	@cd docker && $(MAKE) run
run-test:
	@cd docker && $(MAKE) run-test

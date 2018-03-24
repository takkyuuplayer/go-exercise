.PHONY: test run run-test

setup: dep

dep:
	which dep || go get -u github.com/golang/dep/cmd/dep
	dep ensure
	dep ensure -update

test:
	go test ./...

run:
	@cd docker && $(MAKE) run
run-test:
	@cd docker && $(MAKE) run-test

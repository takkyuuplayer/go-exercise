.PHONY: test

setup:
	go mod download

test:
	go clean -testcache
	go test ./...

lint: golint gocyclo

golint:
	which golint || go get -u -v golang.org/x/lint/golint
	go list ./... | xargs ${GOPATH}/bin/golint

gocyclo:
	which gocyclo || go get -u -v github.com/fzipp/gocyclo/cmd/gocyclo
	${GOPATH}/bin/gocyclo -over 20 .

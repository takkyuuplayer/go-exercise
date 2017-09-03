test:
	go test ./...
go:
	@cd docker && $(MAKE) go

go-test:
	@cd docker && $(MAKE) go-test

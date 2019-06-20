.PHONY: all
all: test

.PHONY: clean
clean:
	@go clean -i ./...

.PHONY: test
test:
	@go test -cover -coverprofile ./coverage.out ./...

.PHONY: travis
travis:
	@go test -race -cover -covermode=atomic -coverprofile ./coverage.out ./...

.PHONY: cover
cover: test
	@echo ""
	@go tool cover -func ./coverage.out

.PHONY: bench
bench:
	@go test -benchmem -bench=. ./...


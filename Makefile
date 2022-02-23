COVERAGE_FILE_NAME := coverage.out
GIT_REV := $(shell git rev-parse HEAD)

bench:
	@go test -bench=. -benchmem | tee benchmarks/$(GIT_REV).txt

test:
	@go test -v -cover -coverprofile=$(COVERAGE_FILE_NAME) -covermode=atomic

lint:
	@go vet
	@golint

show-cover-html:
	@go tool cover -html=$(COVERAGE_FILE_NAME)

benchstat:
	@benchstat benchmarks/*

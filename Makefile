PROF_FILE_NAME := prof.out
GIT_REV := $(shell git rev-parse HEAD)

bench:
	@go test -bench=. -benchmem | tee benchmarks/$(GIT_REV).txt

test:
	@go test -v -cover -race

coverprofile:
	@go test -coverprofile=$(PROF_FILE_NAME)

show-html-coverprofile:
	@go tool cover -html=$(PROF_FILE_NAME)

benchstat:
	@benchstat benchmarks/*

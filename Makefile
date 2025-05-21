COUNT ?= 1

all: test

.PHONY: test
test:
	go test -race -v -count=$(COUNT) ./...

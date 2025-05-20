COUNT ?= 1

.PHONY: test
test:
	go test -race -v -count=$(COUNT) ./...

.DEFAULT_GOAL := build

.PHONY: build
build:
	go build -v ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	go clean
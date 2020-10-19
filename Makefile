.PHONY: all
all: test build

.PHONY: build
build:
	go build -o bin/skyfall ./cmd/skyfall/main.go

.PHONY: test
test:
	go test ./...

.PHONY: all
all: generate test

.PHONY: generate
generate:
	docker run -v $(PWD):/opt/protos/src -ti ghcr.io/dane/protos:v0.0.1 generate

.PHONY: test
test:
	go test ./...

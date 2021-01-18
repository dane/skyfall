DOCKER = $(if $(shell which podman), podman, docker)

.PHONY: all
all: protos test

.PHONY: protos
protos:
	${DOCKER} run -v $(PWD):/opt/protos/src --privileged -ti ghcr.io/dane/protos:v0.0.1 generate

.PHONY: test
test:
	go test ./... -race

.PHONY: proposal
proposal:
	cp proposals/TEMPLATE.md proposals/$(shell date +%Y%m%d)-$(NAME).md

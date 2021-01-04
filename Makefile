.PHONY: all
all: generate test

.PHONY: generate
generate:
	docker run -v $(PWD):/opt/protos/src -ti ghcr.io/dane/protos:v0.0.1 generate

.PHONY: mocks
mocks:
	docker run -v $(PWD):/app -ti ekofr/gomock mockgen \
	  -package=v1 \
	  -source=/app/service/v1/validator.go \
	  -destination=/app/service/v1/mock_validator.go

.PHONY: test
test:
	go test ./... -race

.PHONY: proposal
proposal:
	cp proposals/TEMPLATE.md proposals/$(shell date +%Y%m%d)-$(NAME).md

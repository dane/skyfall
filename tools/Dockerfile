FROM golang:1.15.6-alpine AS build

WORKDIR /go/src/github.com/dane/skyfall/tools

COPY tools.go .

RUN apk add git

RUN go mod init && go mod tidy
RUN go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN mkdir -p /go/src/github.com/{protocolbuffers,googleapis}
RUN git clone --depth 1 -b v3.14.0 https://github.com/protocolbuffers/protobuf /go/src/github.com/protocolbuffers/
RUN git clone --depth 1 -b master https://github.com/googleapis/googleapis.git /go/src/github.com/googleapis

FROM alpine:3.12

WORKDIR /proto

RUN apk add protoc

COPY bin/generate /usr/local/bin/generate

COPY --from=build /go/bin/protoc-gen-grpc-gateway /usr/local/bin/protoc-gen-grpc-gateway
COPY --from=build /go/bin/protoc-gen-openapiv2 /usr/local/bin/protoc-gen-openapiv2
COPY --from=build /go/bin/protoc-gen-go /usr/local/bin/protoc-gen-go
COPY --from=build /go/bin/protoc-gen-go-grpc /usr/local/bin/protoc-gen-go-grpc
COPY --from=build /go/src/github.com/protocolbuffers/src /usr/local/include
COPY --from=build /go/src/github.com/googleapis/google/api /usr/local/include/google/api

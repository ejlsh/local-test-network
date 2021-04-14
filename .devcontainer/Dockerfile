ARG VARIANT="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends protobuf-compiler golang-goprotobuf-dev

RUN go get -x google.golang.org/protobuf/cmd/protoc-gen-go \
              google.golang.org/grpc/cmd/protoc-gen-go-grpc

ARG VERSION=stable-slim

FROM debian:${VERSION} 

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt update && \
    apt -y install build-essential git golang-go

ARG PROJECT=gofib
WORKDIR /workspaces/${PROJECT}

ENV GOPATH=/go
ENV GOBIN=/go/bin
ENV GO111MODULE=on
RUN go get -v golang.org/x/tools/gopls@v0.7.2
RUN go get -v github.com/go-delve/delve/cmd/dlv
RUN go get -v github.com/go-delve/delve/cmd/dlv@v1.7.2
RUN go get -v github.com/cweill/gotests/gotests
COPY fibonacci fibonacci/
COPY gofib.go .
RUN go mod init github.com/jac18281828/gofib
RUN go build

WORKDIR /workspaces/${PROJECT}
CMD ./gofib 1000 11000

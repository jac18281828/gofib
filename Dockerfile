ARG VERSION=102221

FROM jac18281828/godev:${VERSION} 

ARG PROJECT=gofib
WORKDIR /workspaces/${PROJECT}

COPY fibonacci fibonacci/
COPY gofib.go .
RUN go mod init github.com/jac18281828/gofib
RUN go build
RUN go test ./...

WORKDIR /workspaces/${PROJECT}
CMD ./gofib 1000 1100

ARG VERSION=latest
FROM jac18281828/godev:${VERSION} 

ARG PROJECT=gofib
WORKDIR /workspaces/${PROJECT}
ENV GOMAXPROCS=10
COPY fibonacci fibonacci/
COPY gofib.go .
RUN chown -R jac.jac .
USER jac
RUN go mod init github.com/jac18281828/gofib
RUN go build
RUN go test ./...
CMD ./gofib 100 2000

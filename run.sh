#!/usr/bin/env bash

VERSION=$(git rev-parse HEAD | cut -c 1-10)

PROJECT=jac18281828/gofib

docker build . -t ${PROJECT}:${VERSION} && \
    docker run --rm -i -t ${PROJECT}:${VERSION}

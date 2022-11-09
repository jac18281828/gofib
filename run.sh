#!/usr/bin/env bash

VERSION=$(date +%m%d%y)

PROJECT=jac18281828/gofib

docker build . -t ${PROJECT}:${VERSION} && \
	docker run --rm -i -t ${PROJECT}:${VERSION}

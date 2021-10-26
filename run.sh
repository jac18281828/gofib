#!/usr/bin/env bash

VERSION=$(date +%m%d%y)

docker build . -t gofib:${VERSION} && \
	docker run --rm -i -t gofib:${VERSION}

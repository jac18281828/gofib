#!/usr/bin/env bash

VERSION=$(date +%s)

docker build . -t gofib:${VERSION} && \
	docker run --rm -i -t gofib:${VERSION}

#!/usr/bin/env bash

docker build . -t go:1 && docker run --rm -i -t go:1

#!/usr/bin/env bash

docker run --rm -v $(pwd):/code golang:1.15 go run /code/cmd/cli/main.go
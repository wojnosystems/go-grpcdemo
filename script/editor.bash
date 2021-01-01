#!/usr/bin/env bash
ROOT=$(dirname "$1")
docker run --rm -ti --user $(id -u):$(id -g) -v $(pwd)/"$ROOT":/"$ROOT" -v $(pwd)/typewriter:/typewriter cwojno/go-typewriter:latest -- /"$1"

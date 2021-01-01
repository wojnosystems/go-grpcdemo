#!/usr/bin/env bash
ROOT=$(dirname "$1")
docker run --rm -ti -v $(pwd)/"$ROOT":/"$ROOT" -v $(pwd)/typewriter:/typewriter go-typewriter:latest -- /"$1"

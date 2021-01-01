#!/usr/bin/env bash

docker run --rm --user "$(id -u)":"$(id -g)" -v "$(pwd)":/defs namely/protoc-all -f protos/grpcdemo.proto -l go -o /defs

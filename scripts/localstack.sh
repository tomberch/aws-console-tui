#!/usr/bin/env bash

#
#  Docker command to start localstack for testing purposes
#
docker run \
  --rm -it \
  -p 4566:4566 \
  -p 4510-4559:4510-4559 \
  localstack/localstack

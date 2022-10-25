#!/bin/bash
docker build -t ghcr.io/openline-kamailio-server --build-arg ARCH=amd64/ .
